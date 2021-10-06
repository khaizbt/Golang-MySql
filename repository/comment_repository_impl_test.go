package repository

import (
	"context"
	CRUD_MySql "crud_mysql"
	"crud_mysql/entity"
	"fmt"
	"testing"
	"time"
)

func TestCommentRepositoryImpl_Insert(t *testing.T) {
	commentRepository := commentRepositoryImpl{CRUD_MySql.GetConnection()}

	ctx := context.Background()
	comment := entity.Comment{
		UserId:    2,
		PostId:    2,
		Comment:   "G siap ni",
		CreatedAt: time.Now(),
	}

	result, err := commentRepository.Insert(ctx, comment)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestCommentRepositoryImpl_FindById(t *testing.T) {
	commentRepository := commentRepositoryImpl{CRUD_MySql.GetConnection()}

	ctx := context.Background()

	result, err := commentRepository.FindById(ctx, 1)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestCommentRepositoryImpl_FindByPost(t *testing.T) {
	commentRepository := commentRepositoryImpl{CRUD_MySql.GetConnection()}

	ctx := context.Background()

	results, err := commentRepository.FindByPost(ctx, 2)

	if err != nil {
		panic(err)
	}

	for _, result := range results {
		fmt.Println(result)
	}
}

func TestCommentRepositoryImpl_FindByUser(t *testing.T) {
	commentRepository := commentRepositoryImpl{CRUD_MySql.GetConnection()}

	ctx := context.Background()

	results, err := commentRepository.FindByUser(ctx, 2)

	if err != nil {
		panic(err)
	}

	for _, result := range results {
		fmt.Println(result)
	}
}

func TestCommentRepositoryImpl_UpdateComment(t *testing.T) {
	commentRepository := commentRepositoryImpl{CRUD_MySql.GetConnection()}

	ctx := context.Background()
	comment := entity.Comment{
		Comment: "Mana bisa lah",
	}
	result, err := commentRepository.UpdateComment(ctx, comment, 1)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestPostRepositoryImpl_DeleteComment(t *testing.T) {
	commentRepository := NewCommentRepository(CRUD_MySql.GetConnection())

	ctx := context.Background()

	err := commentRepository.DeleteComment(ctx, 1)

	if err != nil {
		panic(err)
	}

	fmt.Println("Sukses Delete Data")
}
