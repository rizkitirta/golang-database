package user

import (
	"context"
	"fmt"
	golangdatabase "golang-database"
	"golang-database/entity"
	"testing"
	_ "github.com/go-sql-driver/mysql"
)

func TestInsertUser(t *testing.T) {
	userRepository := NewUserRepository(golangdatabase.GetConnection())
	user := entity.User{
		Username: "user_2",
		Password: "user_2",
	}
	result,err := userRepository.Insert(context.Background(),user)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	userRepository := NewUserRepository(golangdatabase.GetConnection())
	result,err := userRepository.FindById(context.Background(),59)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
func TestFindAll(t *testing.T) {
	userRepository := NewUserRepository(golangdatabase.GetConnection())
	result,err := userRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, val := range result {
		fmt.Println("=========================")
		fmt.Println("Username: ",val.Username)
		fmt.Println("Password: ",val.Password)
	}
}

func TestUpdateById(t *testing.T)  {
	userRepository := NewUserRepository(golangdatabase.GetConnection())
	user := entity.User{
		Username: "user_update1",
		Password: "user_update1",
	}

	result,err := userRepository.UpdateById(context.Background(),user,59)
	if err != nil {
		panic(err)
	}

	fmt.Println("=========================")
	fmt.Println("Username",result.Username)
	fmt.Println("Password",result.Password)
	fmt.Println("=========================")
}

func TestDeleteById(t *testing.T)  {
	userRepository := NewUserRepository(golangdatabase.GetConnection())

	_,err := userRepository.DeleteById(context.Background(),59) 
	if err != nil {
		panic(err)
	}

	fmt.Println("User deleted")
}