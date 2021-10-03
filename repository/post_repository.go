package repository

import (
	"context"
	"crud_mysql/entity"
)

type (
	PostRepository interface {
		Insert(ctx context.Context, post entity.Post) (entity.Post, error)
		FindById(ctx context.Context, id int32) (entity.Post, error)
		FindAll(ctx context.Context) ([]entity.Post, error)
		UpdatePost(ctx context.Context, post entity.Post, id int32) (entity.Post, error)
		FindPostByUser(ctx context.Context, userId int32) ([]entity.Post, error)
		DeletePost(ctx context.Context, id int32) error
	}
)
