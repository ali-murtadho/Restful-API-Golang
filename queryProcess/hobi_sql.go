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
	table_hobi          = "hobi"
)

// GET ALL
func GetAllHobi(ctx context.Context)([]models.Hobi, error){
	var hobis []models.Hobi

	//CEK KONFIGURASI DB
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Tidak bisa menghubungkan ke mysql", err)
	}

	// QUERY
	query := fmt.Sprintf("SELECT * FROM %v", table_hobi)
	rowQuery, err := db.QueryContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	for rowQuery.Next(){
		var hobi models.Hobi

		if err = rowQuery.Scan(&hobi.Id,
			&hobi.Nama_Hobi);
			err != nil{
				return nil, err
			}

			hobis = append(hobis, hobi)
		}
		return hobis, nil
	}

func InsertDataHobi(ctx context.Context, hobi models.Hobi) error{
		//CEK KONFIGURASI DB
		db, err := config.MySQL()
		if err != nil {
			log.Fatal("Tidak bisa menghubungkan ke mysql", err)
		}
		query := fmt.Sprintf("INSERT INTO %v (nama_hobi) VALUES ('%v')", table_hobi, hobi.Nama_Hobi)
		_, err = db.ExecContext(ctx, query)

		if err != nil {
			return err
		}

	return nil
}

func UpdateDataHobi(ctx context.Context, hobi models.Hobi, id string) error{
		//CEK KONFIGURASI DB
		db, err := config.MySQL()
		if err != nil {
			log.Fatal("Tidak bisa menghubungkan ke mysql", err)
		}
		query := fmt.Sprintf("UPDATE %v SET nama_hobi = '%v'  WHERE id=%v", table_hobi, hobi.Nama_Hobi, id)
		_, err = db.ExecContext(ctx, query)

		if err != nil {
			return err
		}

	return nil
}

func DeleteDataHobi(ctx context.Context, id string) error{
	//CEK KONFIGURASI DB
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Tidak bisa menghubungkan ke mysql", err)
	}
	query := fmt.Sprintf("DELETE FROM %v WHERE id=%v", table_hobi, id)
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