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

func BenchmarkSplitRawLine(t *testing.B) {
	raw := []byte(`494NML-872/UA1JTC-BQOLY1990-06261.81IDR393.38095
	495NML-872/UA1TUI-ZMYBU1990-06204.53IDR493.38095
	496NML-872/UA1RSH-UKACW1990-06165.48CRC593.38095
	497NML-872/UA1VFS-GARAE1990-06165.48CRC593.38095
	498NML-872/UA1UQQ-GTAJG1990-06210.1CRC593.38095
	499NML-872/UA1MMQ-SUTHG1990-06210.1CRC593.38095
	500NML-872/UA1HAE-OBFTM1990-06220.85CRC593.38095
	501NML-872/UA1OUH-VVLYT1990-06290.97CRC53.21027
	502NML-872/UA1SBH-TNQSC1990-06290.97CRC53.21027
	503NML-872/UA1NHH-XQNEI1990-06290.97CRC53.21027
	504NML-872/UA1ORT-WWXWP1990-06290.97KMF324.37307
	505NML-872/UA1QOL-JYALA1990-06290.97KMF424.37307
	506NML-872/UA1WBY-YGJGZ1990-06290.97KMF424.37307
	507NML-872/UA1OMW-TKXQO1990-06290.97KMF41.13812
	508NML-872/UA1OUY-YWRCQ1990-06290.97KMF435.24146
	509NML-872/UA1PXR-EDINJ1990-06290.97KMF435.24146
	510NML-872/UA1ZZR-AVRCX1990-06267.72KMF472.50765
	511NML-872/UA1ZTI-YXBPF1990-06267.72IRR472.50765
	512NML-872/UA1ZVI-THDAD1990-06122.09IRR472.50765
	513NML-872/UA1GZY-PYCNL1990-06122.09DJF472.50765
	514NML-872/UA1AZH-GPUCR1990-06122.09DJF472.50765
	515NML-872/UA1DGC-XKOHV1990-06122.09DJF472.50765
	516NML-872/UA1RLM-ARTYS1990-06199.49DJF456.1928
	517NML-872/UA1GDG-JSPGE1990-06199.49DJF456.1928
	518NML-872/UA1CXP-UQIIG1990-06199.49DJF456.1928
	519NML-872/UA1FSD-MAWGF1990-06199.49DJF456.1928
	520NML-872/UA1XFL-TXWXZ1990-06168.91DJF456.1928
	521NML-872/UA1MGV-TVHRO1990-06168.91DJF456.1928
	522NML-872/UA1XKT-ARZOL1990-06168.91GBP556.1928
	523NML-872/UA1UFX-LUVGH1990-06168.91GBP556.1928
	524NML-872/UA1KWF-EMJHP1990-06172.29GBP556.1928
	525NML-872/UA1ENW-VHIHE1990-06142.35MDL534.10422
	526NML-872/UA1YYY-YOGQI1990-06142.35MDL534.10422
	527NML-872/UA1GSH-QDVUI1990-06142.35MDL534.10422
	528NML-872/UA1BWS-ZIZZQ1990-06180.09HUF534.10422
	529NML-872/UA1YGN-DHNGI1990-06180.09HUF534.10422
	530NML-872/UA1KDG-WEZEE1990-06180.09HUF534.10422
	531NML-872/UA1KGY-UJXIV1990-06180.09HUF534.10422
	532NML-872/UA1LFG-ONAVR1990-06180.09HUF534.10422
	533NML-872/UA1DXY-EFMXQ1990-06180.09HUF534.10422
	534NML-872/UA1SVF-HGEOQ1990-06188.31HUF534.10422
	535NML-872/UA1DTR-YVQWF1990-06122HUF55.37179
	536NML-872/UA1JAT-RHVYU1990-06122HUF55.37179
	537NML-872/UA1XPE-JAVEY1990-06122HUF55.37179
	538NML-872/UA1XMO-DNSXX1990-06122HUF55.37179
	539NML-872/UA1AVQ-ZXCCB1990-06122HUF45.37179
	540NML-872/UA1TNM-GKACZ1990-06282.77HUF461.59652
	541NML-872/UA1JWY-AOLXV1990-06282.77HUF461.59652
	542NML-872/UA1BXR-SMXPW1990-06282.77HUF461.59652
	543NML-872/UA1NAY-IUWLJ1990-06282.77HUF459.19192
	544NML-872/UA1AJQ-VJJCA1990-08282.77AWG359.19192
	545NML-872/UA1UXC-RXQRQ1990-08282.77AWG459.19192
	546NML-872/UA1XRH-FEFLB1990-08282.77MDL559.19192
	547NML-872/UA1JAP-AZIOS1990-08282.77MDL559.19192
	548NML-872/UA1UEZ-FUCDC1990-08203.27MDL559.19192
	549NML-872/UA1BRE-HGPAE1990-08203.27CDF559.19192
	550NML-872/UA1ZMV-BZKNP1990-08203.27LAK559.19192
	551NML-872/UA1YIJ-JPIQE1990-08203.27LAK559.19192
	552NML-872/UA1PBK-TPRMT1990-08203.27ERN599.44839
	553NML-872/UA1SGO-OFMQD1990-08203.27ERN599.44839
	554NML-872/UA1LCI-ITHBB1990-08203.27ERN599.44839
	555NML-872/UA1FXM-RIMKW1990-08236.83ERN599.44839
	556NML-872/UA1PJN-SSHDQ1990-08236.83ERN599.44839
	557NML-872/UA1OUT-LSJSI1990-08236.83ERN599.44839
	558NML-872/UA1NGH-DSJTN1990-08236.83ERN563.66585
	559NML-872/UA1ZCY-DIVCC1990-08236.83ERN563.66585
	560NML-872/UA1XKD-XJXCG1990-08237.51ERN54.80546`)
	t.ResetTimer()

	for i := 0; i < t.N; i++ {
		res := make([][]byte, 0, 300)
		res = SplitRawLine(res, raw, rune(0x0A))
	}

}
