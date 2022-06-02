package golangdatabase

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "INSERT INTO customer (name) VALUES('Tirta2')"
	_, err := db.ExecContext(ctx, query)

	if err != nil {
		panic(err)
	}

	fmt.Println("Data berhasil disimpan")
}
func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "SELECT * FROM customer"
	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var email sql.NullString // null handler with sql.NullString
		var balance int32
		var birthDate sql.NullTime
		var maried bool
		var createdAt time.Time

		err = rows.Scan(&id, &name, &email, &balance, &birthDate, &maried, &createdAt)

		if err != nil {
			panic(err)
		}

		fmt.Println("=========================")
		fmt.Println("Name :", name)
		if email.Valid {
			fmt.Println("Email :", email.String)
		}
		fmt.Println("Balance :", balance)
		if birthDate.Valid {
			fmt.Println("BirthDate :", birthDate.Time)
		}
		fmt.Println("Married :", maried)
		fmt.Println("Crated At :", createdAt)
	}

	fmt.Println("Done")
}

// contoh sql injection
func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	username := "admin"
	password := "admin"
	query := "SELECT * FROM users where username = '"+username+"' AND password = '"+password+"' "
	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var id int
		var username string
		var password string

		err = rows.Scan(&id, &username,&password)
		if err != nil {
			panic(err)
		}

		fmt.Println("=========================")
		fmt.Println("Username :", username)
		fmt.Println("Password :", password)
	}else {
		fmt.Println("User tidak ditemukan!")
	}

}

