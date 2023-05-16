package main

import "fmt"

// Soal 1
type Buah struct {
	Nama, Warna string
	AdBijinya   bool
	harga       int
}

// Soal 2
type segitiga struct {
	alas, tinggi int
}

type persegi struct {
	sisi int
}

type persegiPanjang struct {
	panjang, lebar int
}

// Soal 3
type phone struct {
	name, brand string
	year        int
	colors      []string
}

// Soal 4
type Movie struct {
	Title, Genre   string
	Duration, Year int
}

func main() {
	// Soal 1
	nanas := Buah{"Nanas", "Kuning", false, 9000}
	jeruk := Buah{"Jeruk", "Oranye", true, 8000}
	semangka := Buah{"Semangka", "Hijau & Merah", true, 10000}
	pisang := Buah{"Pisang", "Kuning", false, 5000}

	fmt.Println(nanas)
	fmt.Println(jeruk)
	fmt.Println(semangka)
	fmt.Println(pisang)

	// Soal 2
	bangunSegitiga := segitiga{4, 5}
	bangunSegitiga.luasSegitiga()

	bangunPersegi := persegi{10}
	bangunPersegi.luasPersegi()

	bangunPersegiPanjang := persegiPanjang{4, 5}
	bangunPersegiPanjang.luasPersegiPanjang()

	// Soal 3
	ph := &phone{}
	ph.addColor("red")
	ph.addColor("green")
	ph.addColor("blue")
	fmt.Println(ph)

	// Soal 4
	var dataFilm = []Movie{}

	lotr := Movie{
		Title:    "LOTR",
		Duration: 120,
		Genre:    "action",
		Year:     1999,
	}

	avenger := Movie{
		Title:    "avenger",
		Duration: 120,
		Genre:    "action",
		Year:     2019,
	}

	spiderman := Movie{
		Title:    "spiderman",
		Duration: 120,
		Genre:    "action",
		Year:     2004,
	}

	juon := Movie{
		Title:    "juon",
		Duration: 120,
		Genre:    "horror",
		Year:     2004,
	}

	lotr.tambahDataFilm(&dataFilm)
	avenger.tambahDataFilm(&dataFilm)
	spiderman.tambahDataFilm(&dataFilm)
	juon.tambahDataFilm(&dataFilm)

	for _, data := range dataFilm {
		fmt.Println("title:", data.Title)
		fmt.Println("duration:", data.Duration)
		fmt.Println("genre", data.Genre)
		fmt.Println("year", data.Year)
		fmt.Println("")
	}
}

// Soal 2
func (s segitiga) luasSegitiga() {
	result := float64(s.alas*s.tinggi) / 2
	fmt.Println("Luas segitiga =", result)
}

func (p persegi) luasPersegi() {
	result := p.sisi * p.sisi
	fmt.Println("Luas persegi =", result)
}

func (pp persegiPanjang) luasPersegiPanjang() {
	result := pp.panjang * pp.lebar
	fmt.Println("Luas persegi panjang =", result)
}

// Soal 3
func (p *phone) addColor(color string) {
	p.colors = append(p.colors, color)
}

// Soal 4
func (movie Movie) tambahDataFilm(data *[]Movie) {
	*data = append(*data, movie)
}
