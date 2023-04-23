// PADA FILE INI, DIBUAT FUNGSI FUNGSI YANG MENUNJANG PROSES SERVER DARI DATA JURUSAN.
// DIMULAI DARI MENDAPATKAN DATA, MENAMBAH DATA, MENGUBAH DATA, DAN MENGHAPUS DATA JURUSAN.
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

// PROSES FUNGSI MENDAPATKAN DATA JURUSAN
func GetJurusan(rw http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	
	hobis, err := queryProcess.GetAllJurusan(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(rw, hobis, http.StatusOK)
}
// PROSES FUNGSI TAMBAH DATA JURUSAN
func PostJurusan(rw http.ResponseWriter, r *http.Request, _ httprouter.Params){
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(rw, "Gunakan content JSON", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//PROSES INISIALISASI HB SEBAGAI DARI DATA HOBI
	var jur models.Jurusan

	if err := json.NewDecoder(r.Body).Decode(&jur);
	err != nil {
		utils.ResponseJSON(rw, err, http.StatusBadRequest)
		return
	}

	if err := queryProcess.InsertDataJurusan(ctx, jur);
	err != nil  {
		utils.ResponseJSON(rw, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"Status":"Successfully",
	}

	utils.ResponseJSON(rw, res, http.StatusOK)
}

// PROSES FUNGSI UPDATE DATA JURUSAN
func UpdateJurusan(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content JSON", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var jur models.Jurusan

	if err := json.NewDecoder(r.Body).Decode(&jur);
	err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var idJur = ps.ByName("id")

	if err := queryProcess.UpdateDataJurusan(ctx, jur, idJur);
	err != nil  {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"Status":"Successfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

// PROSES FUNGSI HAPUS DATA HOBI
func DeleteJurusan(w http.ResponseWriter, r *http.Request, ps httprouter.Params){

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idJur = ps.ByName("id")

	if err := queryProcess.DeleteDataJurusan(ctx, idJur);
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