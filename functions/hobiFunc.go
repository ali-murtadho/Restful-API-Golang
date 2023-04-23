// PADA FILE INI, DIBUAT FUNGSI FUNGSI YANG MENUNJANG PROSES SERVER DARI DATA HOBI.
// DIMULAI DARI MENDAPATKAN DATA, MENAMBAH DATA, MENGUBAH DATA, DAN MENGHAPUS DATA HOBI.
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

// PROSES FUNGSI MENDAPATKAN DATA HOBI
func GetHobi(rw http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// FUNGSI METHOD GetAllHobi DIDAPAT DARI FOLDER queryProcess 
	// YANG MANA METHOD TERSEBUT BERISI QUERY YANG BERISI PERINTAH UNTUK MENDAPATKAN DATA
	hobis, err := queryProcess.GetAllHobi(ctx)

	if err != nil {
		fmt.Println(err)
	}
	// STATUS OK (200) DIGUNAKAN KETIKA SERVER MENERIMA PESAN SUKSES DARI PESAN FUNGSI MENDAPATKAN DATA
	utils.ResponseJSON(rw, hobis, http.StatusOK)
}
// PROSES FUNGSI TAMBAH DATA HOBI
func PostHobi(rw http.ResponseWriter, r *http.Request, _ httprouter.Params){
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(rw, "Gunakan content JSON", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//PROSES INISIALISASI HB SEBAGAI DARI DATA HOBI
	var hb models.Hobi

	if err := json.NewDecoder(r.Body).Decode(&hb);
	err != nil {
		// StatusBadRequest(400) DIGUNAKAN KETIKA PROSES SERVER GAGAL MENERIMA RESPONSE
		utils.ResponseJSON(rw, err, http.StatusBadRequest)
		return
	}

	if err := queryProcess.InsertDataHobi(ctx, hb);
	err != nil  {
		// StatusInternalServerError (500) DIGUNAKAN KETIKA PROSES RESPONSE GAGAL
		utils.ResponseJSON(rw, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"Status":"Successfully",
	}

	utils.ResponseJSON(rw, res, http.StatusOK)
}
// PROSES FUNGSI UPDATE DATA HOBI
func UpdateHobi(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content JSON", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var hb models.Hobi

	if err := json.NewDecoder(r.Body).Decode(&hb);
	err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	// PROSES INISIALISASI ID PADA DATA HOBI YANG DIBERIKAN KEPADA VARIAVEL idHobi
	var idHobi = ps.ByName("id")

	// METHOD UpdateDataHobi DIAMBIL DARI FOLDER queryProcess UNTUK QUERY PENGUBAHAN DATA HOBI
	if err := queryProcess.UpdateDataHobi(ctx, hb, idHobi);
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
func DeleteHobi(w http.ResponseWriter, r *http.Request, ps httprouter.Params){

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idHobi = ps.ByName("id")

	// METHOD DeleteDataHobi YANG DIAMBIL DARI FOLDER queryProcess DIGUNAKAN UNTUK PENEMPATAN QUERY UNTUK PENGHAPUSAN DATA
	if err := queryProcess.DeleteDataHobi(ctx, idHobi);
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