package main

import "fmt"

func main() {
	// Soal 1
	var luasLigkaran float64
	var kelilingLingkaran float64
	var rad float64

	rad = 7
	count(&kelilingLingkaran, &luasLigkaran, rad)
	fmt.Println("Luas = ", luasLigkaran, " Keliling = ", kelilingLingkaran)

	// Soal 2
	var sentence string
	introduce(&sentence, "John", "laki-laki", "penulis", "30")

	fmt.Println(sentence) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"
	introduce(&sentence, "Sarah", "perempuan", "model", "28")

	fmt.Println(sentence) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	// Soal 3
	var buah = []string{}
	buahFavorit(&buah, "Jeruk")
	buahFavorit(&buah, "Semangka")
	buahFavorit(&buah, "Mangga")
	buahFavorit(&buah, "Strawberry")
	buahFavorit(&buah, "Durian")
	buahFavorit(&buah, "Manggis")
	buahFavorit(&buah, "Alpukat")

	for i, name := range buah {
		fmt.Printf("%d. %s\n", i, name)
	}

	// Soal 4
	var dataFilm = []map[string]string{}

	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	for _, data := range dataFilm {
		for i, value := range data {
			fmt.Println(i, ":", value)
		}
		fmt.Println("")
	}
}

// Soal 1
func count(keliling *float64, luas *float64, rad float64) {
	*keliling = 3.14 * 2 * rad
	*luas = 3.14 * rad * rad
}

// Soal 2
func introduce(sentence *string, name, gender, occupation, age string) {
	var ab string
	if gender == "laki-laki" {
		ab = "Pak"
	} else {
		ab = "Bu"
	}
	*sentence = ab + " " + name + " adalah seorang " + occupation + " yang berusia " + age
}

// Soal 3
func buahFavorit(list *[]string, name string) {
	*list = append(*list, name)
}

// Soal 4
func tambahDataFilm(title, duration, genre, year string, data *[]map[string]string) {
	var temp = map[string]string{
		"title":    title,
		"duration": duration,
		"genre":    genre,
		"year":     year,
	}
	*data = append(*data, temp)
}
