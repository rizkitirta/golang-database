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
