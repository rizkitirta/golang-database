package repository

import (
	"context"
	"fmt"
	golang_database "golang-database"
	"golang-database/entity"
	"testing"
	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(golang_database.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email: "tirta@gmail.com",
		Comment: "test comment 4",
	}
	result,err := commentRepository.Insert(ctx,comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	fmt.Println("Done")
}

func TestCommentFindById(t *testing.T)  {
	commentRepo := NewCommentRepository(golang_database.GetConnection())

	ctx := context.Background()
	comment,err := commentRepo.FindById(ctx,5)
	if err != nil {
		panic(err)
	}

	fmt.Println("=========================")
	fmt.Println("Email: ",comment.Email)
	fmt.Println("Comment: ",comment.Comment)
}

func TestCommentFindAll(t *testing.T)  {
	commentRepo := NewCommentRepository(golang_database.GetConnection())
	result,err := commentRepo.FindAll(context.Background())

	if err != nil {
		panic(err)
	}

	for _,comment := range result {
		fmt.Println("=========================")
		fmt.Println("Email: ",comment.Email)
		fmt.Println("Comment: ",comment.Comment)
	}
}