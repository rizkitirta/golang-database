package golangdatabase

import (
	"context"
	"fmt"
	"testing"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "INSERT INTO customer (name) VALUES('Tirta2')"
	_,err := db.ExecContext(ctx,query)

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
	rows,err := db.QueryContext(ctx,query)
	
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int 
		var name string
		err = rows.Scan(&id,&name)

		if err != nil {
			panic(err)
		}

		fmt.Println(id,name)
	}

	fmt.Println("Done")
}
