package main

import (
	"bytes"
	"testing"
)

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

func TestSplitRawLine(t *testing.T) {
	oneToZero := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	oneToNine := oneToZero[:9]
	empty := oneToZero[0:0]
	h0 := `አማርኛ 
	اردو
	العربية
	تۆرکجه
	فارسی
	مصرى
	پښتو
	文言 	日本語 	한국어 	Авар 	Адыгабзэ 	Аҧсшәа 	Беларуская 	Български 	Македонски 	Српски / srpski
	Українська 	Ελληνικά  	Afrikaans  	Akan  	Alemannisch  	Aragonés 	Armãneashti
	Arpetan  	Asturianu 	Atikamekw 	Avañe'ẽ 	Aymar aru 	Azərbaycanca 	Bahasa Indonesia 	Basa Jawa
	Basa Sunda 	Bosanski  	Català 	Dansk  	Davvisámegiella 	Deitsch 	Deutsch 	Eesti 	English 	Español 	Esperanto
	Euskara 	Français 	Frysk 	Føroyskt 	Gaeilge 	Galego 	Gàidhlig 	Hausa 	Hornjoserbsce 	Hrvatski 	Interlingua
	IsiXhosa 	IsiZulu 	Italiano 	Kreyòl ayisyen  	Latina 	Latviešu 	Lietuvių 	Magyar 	Malti 	Nederlands 	Nedersaksies
	Nordfriisk 	Norsk 	Norsk nynorsk 	Occitan 	Oʻzbekcha/ўзбекча  	Patois 	Plattdüütsch 	Polski 	Português 	Pälzisch 	Română
	Scots 	Shqip  	Slovenčina 	Slovenščina  	Soomaaliga 	Suomi 	Svenska 	Tagalog 	Tiếng Việt 	Türkçe 	West-Vlams 	Ænglisc
	Íslenska 	Čeština  	עברית 	ܐܪܡܝܐ 	नेपाली 	भोजपुरी 	मराठी 	संस्कृतम् 	हिन्दी 	অসমীয়া 	বাংলা 	ਪੰਜਾਬੀ 	ગુજરાતી 	தமிழ் 	తెలుగు
	ಕನ್ನಡ 	සිංහල 	ไทย 	Հայերեն 	ქართული`
	h1 := []byte(h0 + string([]byte{0x1F}) + h0)

	tests := []struct {
		name      string
		data      []byte
		delimiter rune
		count     int
		expected  [][]byte
	}{
		{name: "5 field case",
			data:      append(append(append(oneToZero, oneToZero...), oneToZero...), oneToZero...),
			delimiter: rune(0),
			count:     5,
			expected:  append(append(append(append(append(make([][]byte, 0, 4), oneToNine), oneToNine), oneToNine), oneToNine), empty),
		},
		{name: "utf-8",
			data:      h1,
			delimiter: rune(0x1F),
			count:     2,
			expected:  append(append(make([][]byte, 0, 2), []byte(h0)), []byte(h0)),
		},
	}
	t.Log("Testing split functionality")

	for _, test := range tests {
		sub := func(t *testing.T) {
			result := SplitRawLine(make([][]byte, 0, 10), test.data, test.delimiter)

			c := len(result)
			if c == test.count {
				t.Logf("Count: got %v, expected %v", c, test.count)
			} else {
				t.Errorf("Count: got %v, expected %v", c, test.count)
			}

			for index := range result {
				if bytes.Equal(result[index], test.expected[index]) {
					t.Logf("Contents:  index %v, got %v, expected %v", index, result[index], test.expected[index])
				} else {
					t.Errorf("Contents:  index %v, got %v, expected %v", index, result[index], test.expected[index])
				}
			}
		}
		t.Run(test.name, sub)
	}

}
