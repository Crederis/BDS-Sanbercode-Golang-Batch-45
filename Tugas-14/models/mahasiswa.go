package models

import "time"

type NilaiMahasiswa struct {
	Nama        string    `json:"nama"`
	MataKuliah  string    `json:"mata_kuliah"`
	IndeksNilai string    `json:"indeks_nilai"`
	Nilai       float64   `json:"nilai"`
	ID          int       `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
