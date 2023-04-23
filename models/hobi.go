// Berikut ini merupakan rancangan database dari tabel hobi
package models

type Hobi struct {
	Id        int    `json:"id"` // id sebagai primary key tabel hobi 
	Nama_Hobi string `json:"nama_hobi"`
}