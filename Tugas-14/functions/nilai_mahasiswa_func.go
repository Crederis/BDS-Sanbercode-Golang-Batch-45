package functions

import (
	"Tugas-14/models"
	"Tugas-14/query"
	"Tugas-14/utils"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func getIndeksNilai(nilai uint) string {
	switch {
	case nilai >= 80:
		return "A"
	case nilai >= 70:
		return "B"
	case nilai >= 60:
		return "C"
	case nilai >= 50:
		return "D"
	default:
		return "E"
	}
}

func GetNilai(rw http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	nilai, err := query.GetAllNilai(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(rw, nilai, http.StatusOK)
}

func PostNilai(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var nilai models.NilaiMahasiswa
	if err := json.NewDecoder(r.Body).Decode(&nilai); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if nilai.Nilai > 100 {
		http.Error(w, "Nilai tidak boleh lebih dari 100", http.StatusBadRequest)
	}

	nilai.IndeksNilai = getIndeksNilai(uint(nilai.Nilai))

	if err := query.InsertNilai(ctx, nilai); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}
	utils.ResponseJSON(w, res, http.StatusCreated)
}

func UpdateNilai(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var nilai models.NilaiMahasiswa
	if err := json.NewDecoder(r.Body).Decode(&nilai); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var idNilai = ps.ByName("id")

	if nilai.Nilai > 100 {
		http.Error(w, "Nilai tidak boleh lebih dari 100", http.StatusBadRequest)
	}

	nilai.IndeksNilai = getIndeksNilai(uint(nilai.Nilai))

	if err := query.UpdateNilai(ctx, nilai, idNilai); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}
	utils.ResponseJSON(w, res, http.StatusCreated)
}

func DeleteNilai(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idNilai = ps.ByName("id")
	if err := query.DeleteNilai(ctx, idNilai); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}
	utils.ResponseJSON(w, res, http.StatusCreated)
}
