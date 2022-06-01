package golangdatabase

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/golang_database")

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10) //Set maximum coneksi yang digunakan atau tidak digunakan
	db.SetMaxOpenConns(100) //set maximum koneksi
	db.SetConnMaxIdleTime(5 * time.Minute) //set waktu koneksi yang akan ditutup ketika tidak digunakan
	db.SetConnMaxLifetime(60 * time.Minute) //membuat atau memperbarui koneksi sesuai waktu koneksi yang ditetapkan

	return db
}