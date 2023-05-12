package main

import "fmt"

func main() {
	// Soal 1
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volume)

	//  Soal 2
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john)

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah)

	// Soal 3
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}

	var buahFavoritJohn = buahFavorit("John", buah...)

	fmt.Println(buahFavoritJohn)

	// Soal 4
	var dataFilm = []map[string]string{}
	// buatlah closure function disini
	var tambahDataFilm = func(data ...string) {
		var temp = map[string]string{
			"genre": data[2],
			"jam":   data[1],
			"tahun": data[3],
			"title": data[0],
		}
		dataFilm = append(dataFilm, temp)
	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}
}

// Soal 1
func luasPersegiPanjang(panjang int, lebar int) int {
	return panjang * lebar
}

func kelilingPersegiPanjang(panjang int, lebar int) int {
	return 2 * (panjang + lebar)
}

func volumeBalok(panjang int, lebar int, tinggi int) int {
	return panjang * lebar * tinggi
}

// Soal 2
func introduce(name, gender, occupation, age string) string {
	var ab string

	if gender == "laki-laki" {
		ab = "Pak"
	} else {
		ab = "Bu"
	}

	result := ab + " " + name + " adalah seorang " + occupation + " yang berusia " + age + " tahun"

	return result
}

// Soal 3
func buahFavorit(name string, buah ...string) string {
	// halo nama saya john dan buah favorit saya adalah "semangka", "jeruk", "melon", "pepaya"
	result := "halo nama saya " + name + " dan buah favorit saya adalah " + `"` + buah[0] + `", ` + `"` + buah[1] + `", ` + `"` + buah[2] + `", ` + `"` + buah[3] + `"`

	return result
}
