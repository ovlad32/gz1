package main

import "testing"

var gl int

func BenchmarkTx0(t *testing.B) {
	for i := 0; i < t.N; i++ {
		gl = tx0()
	}
}
func BenchmarkTx1(t *testing.B) {
	for i := 0; i < t.N; i++ {
		gl = tx1()
	}
}
func BenchmarkTx2(t *testing.B) {
	for i := 0; i < t.N; i++ {
		gl = tx2()
	}
}
