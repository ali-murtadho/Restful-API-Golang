// PADA FILE INI, DIBUAT FUNGSI FUNGSI YANG MENUNJANG PROSES SERVER DARI DATA MAHASISWA.
// DIMULAI DARI MENDAPATKAN DATA, MENAMBAH DATA, MENGUBAH DATA, DAN MENGHAPUS DATA MAHASISWA.
// STATUS OK (200) DIGUNAKAN KETIKA SERVER MENERIMA PESAN SUKSES DARI PESAN FUNGSI MENDAPATKAN DATA
// StatusInternalServerError (500) DIGUNAKAN KETIKA PROSES RESPONSE GAGAL KARENA SERVER
// StatusBadRequest(400) DIGUNAKAN KETIKA PROSES KESALAHAN PENGIRIMAN DATA YANG TIDAK SESUAI

package functions

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"rest-api-jobhun-tech-test/models"
	"rest-api-jobhun-tech-test/queryProcess"
	"rest-api-jobhun-tech-test/utils"

	"github.com/julienschmidt/httprouter"
)

// PROSES FUNGSI MENDAPATKAN DATA MAHASISWA
func GetMahasiswa(rw http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	
	// METHOD GetAll YANG DIAMBIL DARI FOLDER queryProcess DIGUNAKAN UNTUK MENEMPATKAN QUERY UNTUK MENGAMBIL DATA MAHASISWA
	mahasiswas, err := queryProcess.GetAll(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(rw, mahasiswas, http.StatusOK)
}
// PROSES FUNGSI TAMBAH DATA MAHASISWA
func PostMahasiswa(rw http.ResponseWriter, r *http.Request, _ httprouter.Params){
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(rw, "Gunakan content JSON", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var mhs models.Mahasiswa

	if err := json.NewDecoder(r.Body).Decode(&mhs);
	err != nil {
		utils.ResponseJSON(rw, err, http.StatusBadRequest)
		return
	}

	// METHOD InsertData YANG DIAMBIL DARI FOLDER queryProcess DIGUNAKAN UNTUK MENEMPATKAN QUERY UNTUK MENAMBAH DATA MAHASISWA
	if err := queryProcess.InsertData(ctx, mhs);
	err != nil  {
		utils.ResponseJSON(rw, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"Status":"Successfully",
	}

	utils.ResponseJSON(rw, res, http.StatusOK)
}

// PROSES FUNGSI UPDATE DATA MAHASISWA
func UpdateMahasiswa(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content JSON", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var mhs models.Mahasiswa

	if err := json.NewDecoder(r.Body).Decode(&mhs);
	err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}
	// INSIALISASI id PADA MAHASISWA DENGAN PEMBERIAN VARIABEL idMhs
	var idMhs = ps.ByName("id")

	// METHOD UpdateData YANG DIAMBIL DARI FOLDER queryProcess DIGUNAKAN UNTUK MENEMPATKAN QUERY UNTUK MENGUBAH DATA MAHASISWA
	if err := queryProcess.UpdateData(ctx, mhs, idMhs);
	err != nil  {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"Status":"Successfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

// PROSES FUNGSI HAPUS DATA MAHASISWA
func DeleteMahasiswa(w http.ResponseWriter, r *http.Request, ps httprouter.Params){

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idMhs = ps.ByName("id")

	// METHOD DeleteData YANG DIAMBIL DARI FOLDER queryProcess DIGUNAKAN UNTUK MENEMPATKAN QUERY UNTUK MENGHAPUS DATA MAHASISWA
	if err := queryProcess.DeleteData(ctx, idMhs);
	err != nil  {
		errDel := map[string]string{
			"error" : fmt.Sprintf("%v", err),
		}
		utils.ResponseJSON(w, errDel, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"Status":"Successfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}