package queryProcess

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"rest-api-jobhun-tech-test/config"
	"rest-api-jobhun-tech-test/models"
)

const (
	table_jurusan          = "jurusan"
)

// GET ALL
func GetAllJurusan(ctx context.Context)([]models.Jurusan, error){
	var jurus []models.Jurusan

	//CEK KONFIGURASI DB
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Tidak bisa menghubungkan ke mysql", err)
	}

	// QUERY
	query := fmt.Sprintf("SELECT * FROM %v", table_jurusan)
	rowQuery, err := db.QueryContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	for rowQuery.Next(){
		var jur models.Jurusan

		if err = rowQuery.Scan(&jur.Id,
			&jur.Nama_Jurusan);
			err != nil{
				return nil, err
			}

			jurus = append(jurus, jur)
		}
		return jurus, nil
	}

func InsertDataJurusan(ctx context.Context, jurusan models.Jurusan) error{
		//CEK KONFIGURASI DB
		db, err := config.MySQL()
		if err != nil {
			log.Fatal("Tidak bisa menghubungkan ke mysql", err)
		}
		query := fmt.Sprintf("INSERT INTO %v (nama_jurusan) VALUES ('%v')", table_jurusan, jurusan.Nama_Jurusan)
		_, err = db.ExecContext(ctx, query)

		if err != nil {
			return err
		}

	return nil
}

func UpdateDataJurusan(ctx context.Context, jurusan models.Jurusan, id string) error{
		//CEK KONFIGURASI DB
		db, err := config.MySQL()
		if err != nil {
			log.Fatal("Tidak bisa menghubungkan ke mysql", err)
		}
		query := fmt.Sprintf("UPDATE %v SET nama_jurusan = '%v'  WHERE id=%v", table_jurusan, jurusan.Nama_Jurusan, id)
		_, err = db.ExecContext(ctx, query)

		if err != nil {
			return err
		}

	return nil
}

func DeleteDataJurusan(ctx context.Context, id string) error{
	//CEK KONFIGURASI DB
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Tidak bisa menghubungkan ke mysql", err)
	}
	query := fmt.Sprintf("DELETE FROM %v WHERE id=%v", table_jurusan, id)
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