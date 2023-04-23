package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// INISIALISASI DATABASE YANG AKAN DIGUNAKAN, DISINI SAYA MENGGUNAKAN DATABASE MYSQL
const (
	username string = "root"
	password string = ""
	database string = "db_mahasiswa"
)

var dsn = fmt.Sprintf("%v:%v@/%v", username, password, database)

// dsn = "root:password@tcp(localhost)/db_mahasiswa"

func MySQL()(*sql.DB, error){
	db, err := sql.Open("mysql",dsn)

	if err != nil {
		return nil, err
	}
	return db, nil
}
