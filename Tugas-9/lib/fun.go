package lib

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Soal 1
type SegitigaSamaSisi struct {
	Alas, Tinggi int
}

type PersegiPanjang struct {
	Panjang, Lebar int
}

type Tabung struct {
	JariJari, Tinggi float64
}

type Balok struct {
	Panjang, Lebar, Tinggi int
}

type HitungBangunDatar interface {
	Luas() int
	Keliling() int
}

type HitungBangunRuang interface {
	Volume() float64
	LuasPermukaan() float64
}

func (s SegitigaSamaSisi) Luas() int {
	return (s.Alas * s.Tinggi) / 2
}

func (s SegitigaSamaSisi) Keliling() int {
	return s.Alas * 3
}

func (p PersegiPanjang) Luas() int {
	return p.Panjang * p.Lebar
}

func (p PersegiPanjang) Keliling() int {
	return 2 * (p.Panjang + p.Lebar)
}

func (t Tabung) Volume() float64 {
	return math.Pi * math.Pow(t.JariJari, 2) * t.Tinggi
}

func (t Tabung) LuasPermukaan() float64 {
	return 2 * math.Pi * t.JariJari * (t.JariJari + t.Tinggi)
}

func (b Balok) Volume() float64 {
	return float64(b.Panjang * b.Lebar * b.Tinggi)
}

func (b Balok) LuasPermukaan() float64 {
	return 2 * float64((b.Panjang*b.Lebar)+(b.Lebar*b.Tinggi)+(b.Panjang*b.Tinggi))
}

// Soal 2

type Phone struct {
	Name, Brand string
	Year        int
	Colors      []string
}

type Gadget interface {
	ShowName() string
	ShowBrand() string
	ShowYear() string
	ShowColors() string
}

func (p Phone) ShowName() string {
	return "name: " + p.Name
}

func (p Phone) ShowBrand() string {
	return "brand: " + p.Brand
}

func (p Phone) ShowYear() string {
	return "year: " + strconv.Itoa(p.Year)
}

func (p Phone) ShowColors() string {
	return "colors: " + strings.Join(p.Colors, ", ")
}

func ShowGadgetData(g Gadget) {
	fmt.Println(g.ShowName())
	fmt.Println(g.ShowBrand())
	fmt.Println(g.ShowYear())
	fmt.Println(g.ShowColors())
}

// Soal 3
func LuasPersegi(value int, status bool) interface{} {
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
