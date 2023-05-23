package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"sync"
	"time"
)

// Soal 1
func showPhone() {
	var phones = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}
	group := sync.WaitGroup{}
	sort.Strings(phones)

	for i, data := range phones {
		go printPhone(i+1, data, &group)
		time.Sleep(1 * time.Second)
	}

	group.Wait()
}

func printPhone(index int, phone string, group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	fmt.Printf("%d. %s\n", index, phone)
}

// Soal 2
func getMovies(moviesChannel chan<- string, movies ...string) {
	defer close(moviesChannel)
	for i, data := range movies {
		moviesChannel <- strconv.Itoa(i+1) + ". " + data
	}
	time.Sleep(5 * time.Second)
}

// Soal 3
func luasLingkaran(result chan<- string, rad float64) {
	defer close(result)

	area := math.Pi * math.Pow(rad, 2)
	result <- "Luas lingkaran dengan jari-jari " + strconv.FormatFloat(rad, 'f', 3, 64) + " adalah " + strconv.FormatFloat(area, 'f', 3, 64)

	time.Sleep(5 * time.Second)
}

func kelilingLingkaran(result chan<- string, rad float64) {
	defer close(result)

	perimeter := math.Pi * rad * 2
	result <- "Keliling lingkaran dengan jari-jari " + strconv.FormatFloat(rad, 'f', 3, 64) + " adalah " + strconv.FormatFloat(perimeter, 'f', 3, 64)

	time.Sleep(5 * time.Second)
}

func volumeTabung(result chan<- string, rad float64, tinggi float64) {
	defer close(result)

	volume := math.Pi * math.Pow(rad, 2) * tinggi
	result <- "Volume tabung dengan jari-jari " + strconv.FormatFloat(rad, 'f', 3, 64) + " dan tinggi " + strconv.FormatFloat(tinggi, 'f', 3, 64) + " adalah " + strconv.FormatFloat(volume, 'f', 3, 64)

	time.Sleep(5 * time.Second)
}

// Soal 4
func luasPersegiPanjang(luas chan<- int, panjang, lebar int) {
	defer close(luas)
	luas <- panjang * lebar
	time.Sleep(5 * time.Second)
}

func kelilingPersegiPanjang(keliling chan<- int, panjang, lebar int) {
	defer close(keliling)
	keliling <- 2 * (panjang + lebar)
	time.Sleep(5 * time.Second)
}

func volumeBalok(volume chan<- int, panjang, lebar, tinggi int) {
	defer close(volume)
	volume <- panjang * lebar * tinggi
	time.Sleep(5 * time.Second)
}

func main() {
	// Soal 1
	showPhone()

	// Soal 2
	fmt.Println("")
	var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}

	moviesChannel := make(chan string)

	go getMovies(moviesChannel, movies...)

	fmt.Println("List Movies: ")
	for value := range moviesChannel {
		fmt.Println(value)
	}

	// Soal 3
	fmt.Println("")
	result := make(chan string)

	go luasLingkaran(result, 8)
	go luasLingkaran(result, 14)
	go luasLingkaran(result, 20)
	go kelilingLingkaran(result, 8)
	go kelilingLingkaran(result, 14)
	go kelilingLingkaran(result, 20)
	go volumeTabung(result, 8, 10)
	go volumeTabung(result, 14, 10)
	go volumeTabung(result, 20, 10)

	for data := range result {
		fmt.Println(data)
	}

	// Soal 4
	fmt.Println("")
	luas := make(chan int)
	keliling := make(chan int)
	volume := make(chan int)

	go luasPersegiPanjang(luas, 10, 8)
	go luasPersegiPanjang(luas, 20, 10)
	go luasPersegiPanjang(luas, 15, 12)
	go kelilingPersegiPanjang(keliling, 10, 8)
	go kelilingPersegiPanjang(keliling, 20, 10)
	go kelilingPersegiPanjang(keliling, 15, 12)
	go volumeBalok(volume, 10, 8, 5)
	go volumeBalok(volume, 20, 10, 8)
	go volumeBalok(volume, 15, 12, 9)

	counter := 9

	for {
		if counter == 0 {
			break
		}

		select {
		case data := <-luas:
			fmt.Println("Luas persegi adalah", data)
			counter--
		case data := <-keliling:
			fmt.Println("Keliling persegi adalah", data)
			counter--
		case data := <-volume:
			fmt.Println("Volume balok adalah", data)
			counter--
		}
	}
}
