package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Soal 1
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"

	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	var luasPersegiPanjang int
	var kelilingPersegiPanjang int
	var luasSegitiga int

	panjang, _ := strconv.Atoi(panjangPersegiPanjang)
	lebar, _ := strconv.Atoi(lebarPersegiPanjang)
	alas, _ := strconv.Atoi(alasSegitiga)
	tinggi, _ := strconv.Atoi(tinggiSegitiga)

	kelilingPersegiPanjang = 2 * (panjang + lebar)
	luasPersegiPanjang = panjang * lebar
	luasSegitiga = (alas * tinggi) / 2

	fmt.Println("Keliling Persegi Panjang:", kelilingPersegiPanjang)
	fmt.Println("Luas Persegi Panjang:", luasPersegiPanjang)
	fmt.Println("Luas Segitiga:", luasSegitiga)

	// Soal 2
	var nilaiJohn = 80
	var nilaiDoe = 50

	if nilaiJohn >= 80 {
		fmt.Println("Nilai John: A")
	} else if nilaiJohn >= 70 && nilaiJohn < 80 {
		fmt.Println("Nilai John: B")
	} else if nilaiJohn >= 60 && nilaiJohn < 70 {
		fmt.Println("Nilai John: C")
	} else if nilaiJohn >= 50 && nilaiJohn < 60 {
		fmt.Println("Nilai John: D")
	} else {
		fmt.Println("Nilai John: E")
	}

	if nilaiDoe >= 80 {
		fmt.Println("Nilai Doe: A")
	} else if nilaiDoe >= 70 && nilaiDoe < 80 {
		fmt.Println("Nilai Doe: B")
	} else if nilaiDoe >= 60 && nilaiDoe < 70 {
		fmt.Println("Nilai Doe: C")
	} else if nilaiDoe >= 50 && nilaiDoe < 60 {
		fmt.Println("Nilai Doe: D")
	} else {
		fmt.Println("Nilai Doe: E")
	}

	// Soal 3
	tanggal := 20
	bulan := 10
	tahun := 1997
	var strBulan string

	switch bulan {
	case 1:
		strBulan = "Januari"
	case 2:
		strBulan = "Februari"
	case 3:
		strBulan = "Maret"
	case 4:
		strBulan = "April"
	case 5:
		strBulan = "Mei"
	case 6:
		strBulan = "Juni"
	case 7:
		strBulan = "Juli"
	case 8:
		strBulan = "Agustus"
	case 9:
		strBulan = "September"
	case 10:
		strBulan = "Oktober"
	case 11:
		strBulan = "November"
	case 12:
		strBulan = "Desember"
	}

	fmt.Println(tanggal, strBulan, tahun)

	// Soal 4
	if tahun >= 1944 && tahun <= 1964 {
		fmt.Println("Baby boomer")
	} else if tahun >= 1965 && tahun <= 1979 {
		fmt.Println("Generasi X")
	} else if tahun >= 1980 && tahun <= 1994 {
		fmt.Println("Generasi Y (Millenials)")
	} else if tahun >= 1995 && tahun <= 2015 {
		fmt.Println("Generasi Z")
	}
}
