package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-database/entity"
	"strconv"
)

type CommentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(DB *sql.DB) CommentRepository {
	return &CommentRepositoryImpl{DB: DB}
}

func (repository *CommentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	query := "INSERT INTO comments (email, comment) VALUES (?, ?)"
	result, err := repository.DB.ExecContext(ctx, query, comment.Email,comment.Comment)
	if err != nil {
		panic(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	comment.Id = int32(id)
	return comment, nil
}

func (repository *CommentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	query := "SELECT * FROM comments WHERE id = ?"
	result,err := repository.DB.QueryContext(ctx, query, id)
	comment := entity.Comment{}
	if err != nil {
		panic(err)
	}
	defer result.Close()

	if result.Next() {
		err := result.Scan(&comment.Id, &comment.Email, &comment.Comment)
		if err != nil {
			panic(err)
		}

		return comment, nil
	}else {
		return comment, errors.New("Id" + strconv.Itoa(int(id)) + "not found")
	}
}

func (repository *CommentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	query := "SELECT * FROM comments"
	results,err := repository.DB.QueryContext(ctx, query)
	
	if err != nil {
		return nil, err
	}
	defer results.Close()

	comments := []entity.Comment{}
	for results.Next() {
		comment := entity.Comment{}
		err := results.Scan(&comment.Id, &comment.Email, &comment.Comment)
		if err != nil {
			panic(err)
		}

		comments = append(comments, comment)
	}

	return comments, nil
}
