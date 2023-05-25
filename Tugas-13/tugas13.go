package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type NilaiMahasiswa struct {
	Nama, MataKuliah, IndeksNilai string
	Nilai, ID                     uint
}

var nilaiNilaiMahasiswa = []NilaiMahasiswa{}

// Soal 1
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uname, pwd, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte("Username atau Password tidak boleh kosong"))
			return
		}

		if uname == "admin" && pwd == "admin" {
			next.ServeHTTP(w, r)
			return
		}
		w.Write([]byte("Username atau Password tidak sesuai"))
		return
	})
}

func postNilai(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var NM NilaiMahasiswa
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			// parse dari json
			decodeJSON := json.NewDecoder(r.Body)
			NM.ID = uint(len(nilaiNilaiMahasiswa) + 1)
			nilai := NM.Nilai
			var indeks string
			switch {
			case nilai >= 80:
				indeks = "A"
			case nilai >= 70:
				indeks = "B"
			case nilai >= 60:
				indeks = "C"
			case nilai >= 50:
				indeks = "D"
			default:
				indeks = "E"
			}
			NM.IndeksNilai = indeks
			if err := decodeJSON.Decode(&NM); err != nil {
				log.Fatal(err)
			}
		} else {
			// parse dari form
			getNama := r.PostFormValue("nama")
			getMataKuliah := r.PostFormValue("mataKuliah")
			getNilai := r.PostFormValue("nilai")
			nilai, _ := strconv.Atoi(getNilai)
			var indeks string
			switch {
			case nilai >= 80:
				indeks = "A"
			case nilai >= 70:
				indeks = "B"
			case nilai >= 60:
				indeks = "C"
			case nilai >= 50:
				indeks = "D"
			default:
				indeks = "E"
			}
			NM = NilaiMahasiswa{
				Nama:        getNama,
				MataKuliah:  getMataKuliah,
				IndeksNilai: indeks,
				Nilai:       uint(nilai),
				ID:          uint(len(nilaiNilaiMahasiswa) + 1),
			}
		}
	}

	data, _ := json.Marshal(NM)
	w.Write(data)
	nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, NM)
	return
}

// Soal 2
func getNilai(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		dataMovies, err := json.Marshal(nilaiNilaiMahasiswa)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(dataMovies)
		return
	}
	http.Error(w, "ERROR....", http.StatusNotFound)
}

func main() {
	http.Handle("/post_nilai", Auth(http.HandlerFunc(postNilai)))
	http.HandleFunc("/get_nilai", getNilai)
	fmt.Println("server running at http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
