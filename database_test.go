package golangdatabase

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {

}


func TestOpenConnection(t *testing.T)  {
	db, err := sql.Open("mysql","root@tcp(localhost:3306)/golang_database")
	defer db.Close()
	
	if err != nil {
		fmt.Println(err)
		t.Error(err)
	}
}