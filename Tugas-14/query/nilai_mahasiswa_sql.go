package query

import (
	"Tugas-14/config"
	"Tugas-14/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

const (
	table          = "nilai_mahasiswa"
	layoutDateTime = "2006-01-02 15:04:05"
)

func GetAllNilai(ctx context.Context) ([]models.NilaiMahasiswa, error) {
	var scores []models.NilaiMahasiswa
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := "SELECT * FROM " + table + " Order By created_at DESC"
	rowQuery, err := db.QueryContext(ctx, queryText)

	for rowQuery.Next() {
		var score models.NilaiMahasiswa
		var created_at, updated_at string
		if err = rowQuery.Scan(&score.ID,
			&score.Nama,
			&score.MataKuliah,
			&score.Nilai,
			&score.IndeksNilai,
			&created_at,
			&updated_at); err != nil {
			return nil, err
		}

		score.CreatedAt, err = time.Parse(layoutDateTime, created_at)
		if err != nil {
			log.Fatal(err)
		}

		score.UpdatedAt, err = time.Parse(layoutDateTime, created_at)
		if err != nil {
			log.Fatal(err)
		}

		scores = append(scores, score)
	}
	return scores, nil
}

func InsertNilai(ctx context.Context, score models.NilaiMahasiswa) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal(err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (nama, mata_kuliah, nilai, indeks_nilai, created_at, updated_at) values('%v', '%v',%v, '%v', NOW(), NOW())", table,
		score.Nama,
		score.MataKuliah,
		score.Nilai,
		score.IndeksNilai)
	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

func UpdateNilai(ctx context.Context, score models.NilaiMahasiswa, idNilai string) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal(err)
	}

	queryText := fmt.Sprintf("UPDATE %v set nama='%v', mata_kuliah='%v', nilai='%v', indeks_nilai='%v' WHERE id=%v", table,
		score.Nama,
		score.MataKuliah,
		score.Nilai,
		score.IndeksNilai,
		idNilai)
	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

func DeleteNilai(ctx context.Context, idNilai string) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal(err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v WHERE ID=%v", table, idNilai)
	s, err := db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	check, err := s.RowsAffected()

	if check == 0 {
		return errors.New("id tidak ada")
	}

	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}
