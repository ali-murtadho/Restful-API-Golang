package main

import (
	"fmt"
	"log"
	"net/http"
	"rest-api-jobhun-tech-test/functions"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// FUNGSI PENGECEKAN STATUS KONEKSI MYSQL
	// db, err := config.MySQL()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// errorDB := db.Ping()
	// if errorDB != nil {
	// 	panic(errorDB.Error())
	// }
	// fmt.Println("Success")

	router := httprouter.New()

	// ENDPOINT UNTUK PROSES CRUD DATA MAHASISWA
	router.GET("/mahasiswa", functions.GetMahasiswa)
	router.POST("/mahasiswa/insert", functions.PostMahasiswa)
	router.PUT("/mahasiswa/:id/update", functions.UpdateMahasiswa)
	router.DELETE("/mahasiswa/:id/delete", functions.DeleteMahasiswa)
	
	// ENDPOINT UNTUK PROSES CRUD DATA HOBI
	router.GET("/hobi", functions.GetHobi)
	router.POST("/hobi/insert", functions.PostHobi)
	router.PUT("/hobi/:id/update", functions.UpdateHobi)
	router.DELETE("/hobi/:id/delete", functions.DeleteHobi)
	
	// ENDPOINT UNTUK PROSES CRUD DATA JURUSAN
	router.GET("/jurusan", functions.GetJurusan)
	router.POST("/jurusan/insert", functions.PostJurusan)
	router.PUT("/jurusan/:id/update", functions.UpdateJurusan)
	router.DELETE("/jurusan/:id/delete", functions.DeleteJurusan)

	fmt.Println("Server Running at Port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

