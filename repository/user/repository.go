package user

import (
	"context"
	"database/sql"
	"errors"
	"golang-database/entity"
	"strconv"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserInterface {
	return &UserRepository{DB: db}
}

func (repository *UserRepository) Insert(ctx context.Context, user entity.User) (entity.User, error) {
	query := "INSERT INTO users (username,password) VALUES(?,?)"

	result, err := repository.DB.ExecContext(ctx, query, user.Username, user.Password)
	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	user.Id = int32(id)
	return user, nil
}

func (repository *UserRepository) FindById(ctx context.Context, id int32) (entity.User, error) {
	query := "SELECT * FROM users where id = ?"
	result, err := repository.DB.QueryContext(ctx, query, id)
	if err != nil {
		panic(err)
	}
	defer result.Close()

	user := entity.User{}
	if result.Next() {
		err := result.Scan(&user.Id, &user.Username, &user.Password)
		if err != nil {
			panic(err)
		}

		return user, nil
	} else {
		return user, errors.New("User dengan Id " + strconv.Itoa(int(id)) + " tidak ditemukan")
	}
}

func (repository *UserRepository) FindAll(ctx context.Context) ([]entity.User, error) {
	query := "SELECT * FROM users"
	results, err := repository.DB.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer results.Close()

	users := []entity.User{}
	for results.Next() {
		user := entity.User{}
		err := results.Scan(&user.Id, &user.Username, &user.Password)
		if err != nil {
			panic(err)
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository *UserRepository) UpdateById(ctx context.Context, user entity.User,id int32) (entity.User, error) {
	query := "UPDATE users SET username = ?, password=? WHERE id = ?"
	_,err := repository.DB.ExecContext(ctx, query, user.Username,user.Password,id)
	if err != nil {
		panic(err)
	}
	return user,nil
}

func (repository *UserRepository) DeleteById(ctx context.Context, id int32) (entity.User, error) {
	query := "DELETE FROM users WHERE id = ?"
	_,err := repository.DB.ExecContext(ctx, query, id)
	if err != nil {
		panic(err)
	}
	return entity.User{},nil
}
