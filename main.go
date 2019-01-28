package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"runtime"
	"sync"
	"sync/atomic"
	"unsafe"
)

func main() {
	fmt.Println(tx2())
}

func tx0() int {
	f, err := os.Open("C:/home/vlad/data.CK3/data/41/3/ORCL.CK3.TX.dat")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	gf, err := gzip.NewReader(f)
	if err != nil {
		log.Fatal(err)
	}
	defer gf.Close()
	//ba := make([]byte,4096)
	//bytes.NewBuffer(ba);
	r := bufio.NewReader(gf)
	var l int

	for {
		s, err := r.ReadSlice(0xA)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		l = l + len(s)
	}
	return int(l)
}

func tx1() int {
	a := make([]byte, 1000)
	f, err := os.Open("C:/home/vlad/data.CK3/data/41/3/ORCL.CK3.TX.dat")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	gf, err := gzip.NewReader(f)
	if err != nil {
		log.Fatal(err)
	}
	defer gf.Close()

	r := bufio.NewReader(gf)
	var l int

	for {
		s, err := r.ReadSlice(0xA)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		a = a[:len(s)]
		copy(a, s)
		l = l + len(a)
	}
	return int(l)
}

func tx2() int {
	gc := runtime.NumCPU()
	ia := make(chan []byte, gc)
	oa := make(chan []byte, gc)
	for i := 0; i < gc; i++ {
		ia <- make([]byte, 1000)
	}
	f, err := os.Open("C:/home/vlad/data.CK3/data/41/3/ORCL.CK3.TX.dat")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	gf, err := gzip.NewReader(f)
	if err != nil {
		log.Fatal(err)
	}
	defer gf.Close()
	//ba := make([]byte,4096)
	//bytes.NewBuffer(ba);

	r := bufio.NewReader(gf)
	var l int32
	var wg sync.WaitGroup
	wg.Add(gc)
	for i := 0; i < gc; i++ {
		go func() {
			var ll int32
			defer func() {
				atomic.AddInt32(&l, ll)
				wg.Done()
			}()
			for a := range oa {
				ll = ll + int32(len(a))
				ia <- a
			}
		}()
	}

	for {
		s, err := r.ReadSlice(0xA)
		if err != nil {
			if err == io.EOF {
				close(oa)
				break
			}
			log.Fatal(err)
		}
		select {
		case a := <-ia:
			if cap(a) < len(s) {
				a = make([]byte, len(s))
			} else {
				a = a[:len(s)]
			}
			copy(a, s)
			oa <- a
		}
	}
	wg.Wait()
	close(ia)
	return int(l)
}

func BytesToString(s *string, b *[]byte) {
	sth := (*reflect.StringHeader)(unsafe.Pointer(s))
	slh := (*reflect.SliceHeader)(unsafe.Pointer(b))
	sth.Data = slh.Data
	sth.Len = slh.Len
}

//SplitRawLine ...
func SplitRawLine(result [][]byte,
	rawLine []byte,
	delimiter byte,
) [][]byte {
	startPos := 0
	for currPos, cr := range rawLine {
		if cr == delimiter {
			result = append(result, rawLine[startPos:currPos])
			startPos = currPos + 1
		}
	}
	result = append(result, rawLine[startPos:])
	return result
}
