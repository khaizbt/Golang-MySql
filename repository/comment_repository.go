package repository

import (
	"context"
	"crud_mysql/entity"
)

type (
	CommentRepository interface {
		Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error)
		FindById(ctx context.Context, id int32) (entity.Comment, error)
		FindByPost(ctx context.Context, postId int32) ([]entity.Comment, error)
		FindByUser(ctx context.Context, userId int32) ([]entity.Comment, error)
		UpdateComment(ctx context.Context, comment entity.Comment, id int32) (entity.Comment, error)
		DeleteComment(ctx context.Context, id int32) error
	}
)
