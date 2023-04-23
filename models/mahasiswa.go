// Berikut ini merupakan rancangan database dari tabel Mahasiswa
package models

import "time"

type Mahasiswa struct {
	Id                 int    `json:"id"` // id sebagai primary key tabel mahasiswa
	Nama               string `json:"nama"` 
	Usia               int    `json:"usia"`
	Gender             int    `json:"gender"`
	HobiId 			   int 	  `json:"id_hobi"` // kolom id_hobi digunakan sebagai foreign key dari tabel hobi
	JurusanId 		   int 	  `json:"id_jurusan"` // kolom id_jurusan digunakan sebagai foreign key dari tabel jurusan
	Tanggal_Registrasi time.Time `json:"tanggal_registrasi"`
}