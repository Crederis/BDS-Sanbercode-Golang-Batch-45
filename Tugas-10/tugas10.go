package main

import (
	"errors"
	"flag"
	"fmt"
	"math"
	"strconv"
	"time"
)

// Soal 1
func sayTheWord(kalimat string, tahun int) {
	fmt.Println(kalimat, tahun)
}

// Soal 2
func endProcess() {
	message := recover()
	fmt.Println("Terjadi error:", message)
}

func kelilingSegitigaSamaSisi(num int, status bool) (string, error) {
	if status {
		if num == 0 {
			return "", errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
		} else {
			keliling := 3 * num
			result := "keliling segitiga sama sisinya dengan sisi " + strconv.Itoa(num) + " cm adalah " + strconv.Itoa(keliling) + " cm"
			return result, nil
		}
	} else {
		if num == 0 {
			defer endProcess()
			panic("Maaf anda belum menginput sisi dari segitiga sama sisi")
		} else {
			keliling := 3 * num
			return strconv.Itoa(keliling), nil
		}
	}
}

func tambahAngka(num int, angka *int) {
	*angka += num
}

func cetakAngka(angka *int) {
	fmt.Println(*angka)
}

// Soal 4
func tambahPonsel(name string, phones *[]string) {
	*phones = append(*phones, name)
}

// Soal 5
func kelilingLingkaran(rad float64) float64 {
	return math.Round(math.Pi * 2 * rad)
}

func luasLingkaran(rad float64) float64 {
	return math.Round(math.Pi * math.Pow(rad, 2))
}

func main() {
	angka := 1

	// Soal 1
	defer sayTheWord("Golang Backend Development", 2021)
	// fmt.Println("Ini harus jalan duluan")
	// fmt.Println("Ini juga harus jalan duluan")

	// Soal 2
	fmt.Println(kelilingSegitigaSamaSisi(4, true))

	fmt.Println(kelilingSegitigaSamaSisi(8, false))

	fmt.Println(kelilingSegitigaSamaSisi(0, true))

	fmt.Println(kelilingSegitigaSamaSisi(0, false))

	// Soal 3

	defer cetakAngka(&angka)

	tambahAngka(7, &angka)

	tambahAngka(6, &angka)

	tambahAngka(-1, &angka)

	tambahAngka(9, &angka)

	// Soal 4
	var phones = []string{}

	tambahPonsel("Xiaomi", &phones)
	tambahPonsel("Asus", &phones)
	tambahPonsel("iPhone", &phones)
	tambahPonsel("Samsung", &phones)
	tambahPonsel("Oppo", &phones)
	tambahPonsel("Realme", &phones)
	tambahPonsel("Vivo", &phones)

	for i, name := range phones {
		fmt.Printf("%d. %s\n", i+1, name)
		time.Sleep(time.Second * 1)
	}

	// Soal 5
	fmt.Println("Luas lingkaran =", luasLingkaran(7))
	fmt.Println("Keliling lingkaran =", kelilingLingkaran(7))

	fmt.Println("Luas lingkaran =", luasLingkaran(10))
	fmt.Println("Keliling lingkaran =", kelilingLingkaran(10))

	fmt.Println("Luas lingkaran =", luasLingkaran(15))
	fmt.Println("Keliling lingkaran =", kelilingLingkaran(15))

	// Soal 6
	var panjang = flag.Int("panjang", 0, "Masukkan panjang")
	var lebar = flag.Int("lebar", 0, "Masukkan lebar")

	flag.Parse()

	kelilingPersegiPanjang := 2 * (*panjang + *lebar)
	luasPersegiPanjang := *panjang * *lebar

	fmt.Println("Keliling Persegi Panjang =", kelilingPersegiPanjang)
	fmt.Println("Luas Persegi Panjang =", luasPersegiPanjang)
}
