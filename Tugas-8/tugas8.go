package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Soal 1
type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi int
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

func (s segitigaSamaSisi) luas() int {
	return (s.alas * s.tinggi) / 2
}

func (s segitigaSamaSisi) keliling() int {
	return s.alas * 3
}

func (p persegiPanjang) luas() int {
	return p.panjang * p.lebar
}

func (p persegiPanjang) keliling() int {
	return 2 * (p.panjang + p.lebar)
}

func (t tabung) volume() float64 {
	return math.Pi * math.Pow(t.jariJari, 2) * t.tinggi
}

func (t tabung) luasPermukaan() float64 {
	return 2 * math.Pi * t.jariJari * (t.jariJari + t.tinggi)
}

func (b balok) volume() float64 {
	return float64(b.panjang * b.lebar * b.tinggi)
}

func (b balok) luasPermukaan() float64 {
	return 2 * float64((b.panjang*b.lebar)+(b.lebar*b.tinggi)+(b.panjang*b.tinggi))
}

// Soal 2

type phone struct {
	name, brand string
	year        int
	colors      []string
}

type gadget interface {
	showName() string
	showBrand() string
	showYear() string
	showColors() string
}

func (p phone) showName() string {
	return "name: " + p.name
}

func (p phone) showBrand() string {
	return "brand: " + p.brand
}

func (p phone) showYear() string {
	return "year: " + strconv.Itoa(p.year)
}

func (p phone) showColors() string {
	return "colors: " + strings.Join(p.colors, ", ")
}

func showGadgetData(g gadget) {
	fmt.Println(g.showName())
	fmt.Println(g.showBrand())
	fmt.Println(g.showYear())
	fmt.Println(g.showColors())
}

// Soal 3
func luasPersegi(value int, status bool) interface{} {
	if status {
		if value == 0 {
			return "Maaf anda belum menginput sisi dari persegi"
		} else {
			luas := int(math.Pow(float64(value), 2))
			result := "Luas persegi dengan sisi " + strconv.Itoa(value) + " cm adalah " + strconv.Itoa(luas) + " cm"
			return result
		}
	} else {
		if value == 0 {
			return nil
		} else {
			return math.Pow(float64(value), 2)
		}
	}
}

func main() {
	// Soal 1
	var segitiga, persegi hitungBangunDatar
	var tb, blk hitungBangunRuang
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

	segitiga = segitigaSamaSisi{alasSegitiga, tinggiSegitiga}
	fmt.Println("Keliling segitiga =", segitiga.keliling())
	fmt.Println("Luas segitiga =", segitiga.luas())

	persegi = persegiPanjang{panjangPersegiPanjang, lebarPersegiPanjang}
	fmt.Println("Keliling persegi =", persegi.keliling())
	fmt.Println("Luas persegi =", persegi.luas())

	tb = tabung{rad, tinggiTabung}
	fmt.Println("Volume tabung =", tb.volume())
	fmt.Println("Luas permukaan tabung =", tb.luasPermukaan())

	blk = balok{panjangBalok, lebarBalok, tinggiBalok}
	fmt.Println("Volume balok =", blk.volume())
	fmt.Println("Luas permukaan balok =", blk.luasPermukaan())

	// Soal 2
	var ph gadget

	ph = phone{
		name:   "Samsung Galaxy Note 20",
		brand:  "Samsung Galaxy Note 20",
		year:   2020,
		colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}

	// fmt.Println(ph.showName())
	// fmt.Println(ph.showBrand())
	// fmt.Println(ph.showYear())
	// fmt.Println(ph.showColors())
	showGadgetData(ph)

	// Soal 3
	fmt.Println(luasPersegi(4, true))

	fmt.Println(luasPersegi(8, false))

	fmt.Println(luasPersegi(0, true))

	fmt.Println(luasPersegi(0, false))

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
