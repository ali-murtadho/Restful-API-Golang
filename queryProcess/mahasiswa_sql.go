package queryProcess

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"rest-api-jobhun-tech-test/config"
	"rest-api-jobhun-tech-test/models"
	"time"
)

const (
	table          = "mahasiswa"
	layoutDateTime = "2006-01-02 15:04:05"
)

// GET ALL
func GetAll(ctx context.Context)([]models.Mahasiswa, error){
	var mahasiswas []models.Mahasiswa

	//CEK KONFIGURASI DB
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Tidak bisa menghubungkan ke mysql", err)
	}

	// QUERY
	query := fmt.Sprintf("SELECT * FROM %v ORDER BY tanggal_registrasi DESC", table)
	rowQuery, err := db.QueryContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	for rowQuery.Next(){
		var mahasiswa models.Mahasiswa
		var tanggal_registrasi string

		if err = rowQuery.Scan(&mahasiswa.Id,
			&mahasiswa.Nama, 
			&mahasiswa.Usia,
			&mahasiswa.Gender,
			&mahasiswa.HobiId,
			&mahasiswa.JurusanId, 
			&tanggal_registrasi);
			err != nil{
				return nil, err
			}

			mahasiswa.Tanggal_Registrasi, err = time.Parse(layoutDateTime, tanggal_registrasi);
			if err != nil {
				log.Fatal(err)
		
			}

			mahasiswas = append(mahasiswas, mahasiswa)
		}
		return mahasiswas, nil
	}

func InsertData(ctx context.Context, mahasiswa models.Mahasiswa) error{
		//CEK KONFIGURASI DB
		db, err := config.MySQL()
		if err != nil {
			log.Fatal("Tidak bisa menghubungkan ke mysql", err)
		}
		query := fmt.Sprintf("INSERT INTO %v (nama, usia, gender, id_hobi, id_jurusan, tanggal_registrasi) VALUES ('%v',%v,%v,%v,%v, NOW())", table, mahasiswa.Nama, mahasiswa.Usia, mahasiswa.Gender, mahasiswa.HobiId, mahasiswa.JurusanId)
		_, err = db.ExecContext(ctx, query)

		if err != nil {
			return err
		}

	return nil
}

func UpdateData(ctx context.Context, mahasiswa models.Mahasiswa, id string) error{
		//CEK KONFIGURASI DB
		db, err := config.MySQL()
		if err != nil {
			log.Fatal("Tidak bisa menghubungkan ke mysql", err)
		}
		query := fmt.Sprintf("UPDATE %v SET nama = '%v', usia = %v, gender = %v, id_hobi = %v, id_jurusan=%v, tanggal_registrasi = NOW() WHERE id=%v", table, mahasiswa.Nama, mahasiswa.Usia, mahasiswa.Gender, mahasiswa.HobiId, mahasiswa.JurusanId ,id)
		_, err = db.ExecContext(ctx, query)

		if err != nil {
			return err
		}

	return nil
}

func DeleteData(ctx context.Context, id string) error{
	//CEK KONFIGURASI DB
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Tidak bisa menghubungkan ke mysql", err)
	}
	query := fmt.Sprintf("DELETE FROM %v WHERE id=%v", table, id)
	s, err := db.ExecContext(ctx, query)

	if err != nil && err != sql.ErrNoRows{
		return err
	}

	check, err := s.RowsAffected()
	fmt.Println(check)
	if check == 0 {
		return errors.New("id tidak ada")
	}

	if err != nil {
		fmt.Println(err.Error())
	}

return nil
}