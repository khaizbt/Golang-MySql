package repository

import (
	"context"
	"crud_mysql/entity"
)

type (
	PostRepository interface {
		Insert(ctx context.Context, post entity.Post) (entity.Post, error)
		FindById(ctx context.Context, id int32) (entity.Post, error)
	}
)
