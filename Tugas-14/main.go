package main

import (
	"Tugas-14/functions"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// db, e := config.MySQL()

	// if e != nil {
	// 	log.Fatal(e)
	// }

	// eb := db.Ping()
	// if eb != nil {
	// 	panic(eb.Error())
	// }

	// fmt.Println("Success")

	router := httprouter.New()
	router.GET("/nilai-mahasiswa", functions.GetNilai)
	router.POST("/nilai-mahasiswa/create", functions.PostNilai)
	router.PUT("/nilai-mahasiswa/:id/update", functions.UpdateNilai)
	router.DELETE("/nilai-mahasiswa/:id/delete", functions.DeleteNilai)

	fmt.Println("Server Running at Port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
