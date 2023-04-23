// Berikut ini merupakan rancangan database dari tabel Jurusan
package models

type Jurusan struct {
	Id         int    `json:"id"` // id sebagai primary key dari tabel jurusan
	Nama_Jurusan        string `json:"nama_jurusan"`
}