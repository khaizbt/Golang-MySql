package repository

import (
	"context"
	"crud_mysql/entity"
)

type (
	UserRepository interface {
		Insert(ctx context.Context, user entity.User) (entity.User, error)
		FindById(ctx context.Context, id int32) (entity.User, error)
		FindAll(ctx context.Context) ([]entity.User, error)
		DeleteById(ctx context.Context, id int32) error
		UpdateData(ctx context.Context, id int32, user entity.User) (entity.User, error)
	}
)
