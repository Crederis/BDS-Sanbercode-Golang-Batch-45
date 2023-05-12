package main

import "fmt"

func main() {
	// Soal 1
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Println(i, " - Berkualitas")
		} else {
			if i%3 == 0 {
				fmt.Println(i, " - I Love Coding")
			} else {
				fmt.Println(i, " - Santai")
			}
		}
	}

	// Soal 2
	alas := 7
	tinggi := 7

	for i := 0; i < tinggi; i++ {
		for j := 0; j < alas; j++ {
			if j <= i {
				fmt.Print("#")
			}
		}
		fmt.Println("")
	}

	// Soal 3
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}

	fmt.Println(kalimat[2:])

	// Soal 4
	var sayuran = []string{}

	sayuran = append(sayuran, "Bayam")
	sayuran = append(sayuran, "Buncis")
	sayuran = append(sayuran, "Kangkung")
	sayuran = append(sayuran, "Kubis")
	sayuran = append(sayuran, "Seledri")
	sayuran = append(sayuran, "Tauge")
	sayuran = append(sayuran, "Timun")

	for i, name := range sayuran {
		fmt.Printf("%d. %s \n", i, name)
	}

	//  Soal 5
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}

	volume := 1
	for key, value := range satuan {
		fmt.Printf("%s = %d \n", key, value)
		volume *= value
	}
	// volume := satuan["panjang"] * satuan["lebar"] * satuan["tinggi"]
	fmt.Println("Volume balok =", volume)
}
