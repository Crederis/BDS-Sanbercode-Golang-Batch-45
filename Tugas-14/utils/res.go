package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, p interface{}, status int) {
	toByte, err := json.Marshal(p)

	if err != nil {
		http.Error(w, "Terjadi kesalahan", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(toByte))
}
