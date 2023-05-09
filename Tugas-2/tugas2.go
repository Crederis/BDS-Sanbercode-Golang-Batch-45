package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Soal 1
	word1, word2, word3, word4, word5 := "Bootcamp", "Digital", "Skill", "Sanbercode", "Golang"

	fmt.Println(word1, word2, word3, word4, word5)

	// Soal 2
	halo := "Halo Dunia"
	new_halo := strings.Replace(halo, "Dunia", "Golang", 1)
	fmt.Println(new_halo)

	// Soal 3
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"

	fmt.Println(kataPertama, strings.Title(kataKedua), strings.Replace(kataKetiga, "r", "R", -1), strings.ToUpper(kataKeempat))

	// Soal 4
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	angka1, err1 := strconv.Atoi(angkaPertama)
	angka2, err2 := strconv.Atoi(angkaKedua)
	angka3, err3 := strconv.Atoi(angkaKetiga)
	angka4, err4 := strconv.Atoi(angkaKeempat)

	if err1 == nil && err2 == nil && err3 == nil && err4 == nil {
		fmt.Println(angka1 + angka2 + angka3 + angka4)
	}

	// Soal 5
	kalimat := `"halo halo bandung"`
	angka := 2021

	new_kalimat := strings.ReplaceAll(kalimat, "halo", "Hi")
	fmt.Println(new_kalimat, "-", angka)
}
