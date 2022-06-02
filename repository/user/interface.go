package user

import (
	"context"
	"golang-database/entity"
)

type UserInterface interface {
	Insert(ctx context.Context, user entity.User) (entity.User, error)
	FindById(ctx context.Context, id int32) (entity.User, error)
	FindAll(ctx context.Context) ([]entity.User,error)
	UpdateById(ctx context.Context, user entity.User, id int32) (entity.User,error) 
	DeleteById(ctx context.Context, id int32) (entity.User,error)
}
