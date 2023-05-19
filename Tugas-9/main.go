package main

import (
	. "Tugas-9/lib"
	"fmt"
	"strconv"
)

func main() {
	// Soal 1
	var segitiga, persegi HitungBangunDatar
	var tb, blk HitungBangunRuang
	var alasSegitiga, tinggiSegitiga, panjangPersegiPanjang, lebarPersegiPanjang, panjangBalok, lebarBalok, tinggiBalok int
	var rad, tinggiTabung float64

	alasSegitiga = 10
	tinggiSegitiga = 10
	panjangPersegiPanjang = 10
	lebarPersegiPanjang = 8
	panjangBalok = 10
	lebarBalok = 8
	tinggiBalok = 5
	rad = 10
	tinggiTabung = 20

	segitiga = SegitigaSamaSisi{alasSegitiga, tinggiSegitiga}
	fmt.Println("Keliling segitiga =", segitiga.Keliling())
	fmt.Println("Luas segitiga =", segitiga.Luas())

	persegi = PersegiPanjang{panjangPersegiPanjang, lebarPersegiPanjang}
	fmt.Println("Keliling persegi =", persegi.Keliling())
	fmt.Println("Luas persegi =", persegi.Luas())

	tb = Tabung{rad, tinggiTabung}
	fmt.Println("Volume tabung =", tb.Volume())
	fmt.Println("Luas permukaan tabung =", tb.LuasPermukaan())

	blk = Balok{panjangBalok, lebarBalok, tinggiBalok}
	fmt.Println("Volume balok =", blk.Volume())
	fmt.Println("Luas permukaan balok =", blk.LuasPermukaan())

	// Soal 2
	var ph Gadget

	ph = Phone{
		Name:   "Samsung Galaxy Note 20",
		Brand:  "Samsung Galaxy Note 20",
		Year:   2020,
		Colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}

	// fmt.Println(ph.showName())
	// fmt.Println(ph.showBrand())
	// fmt.Println(ph.showYear())
	// fmt.Println(ph.showColors())
	ShowGadgetData(ph)

	// Soal 3
	fmt.Println(LuasPersegi(4, true))

	fmt.Println(LuasPersegi(8, false))

	fmt.Println(LuasPersegi(0, true))

	fmt.Println(LuasPersegi(0, false))

	// Soal 4
	var prefix interface{} = "hasil penjumlahan dari "

	var kumpulanAngkaPertama interface{} = []int{6, 8}

	var kumpulanAngkaKedua interface{} = []int{12, 14}

	listAngkaPertama := kumpulanAngkaPertama.([]int)
	listAngkaKedua := kumpulanAngkaKedua.([]int)

	listAngkaPertama = append(listAngkaPertama, listAngkaKedua...)

	var total int

	for _, i := range listAngkaPertama {
		total += i
	}

	result := prefix.(string) + strconv.Itoa(listAngkaPertama[0]) + " + " + strconv.Itoa(listAngkaPertama[1]) + " + " + strconv.Itoa(listAngkaKedua[0]) + " + " + strconv.Itoa(listAngkaKedua[1]) + " = " + strconv.Itoa(total)
	fmt.Println(result)
}
