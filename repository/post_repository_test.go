package repository

import (
	"context"
	CRUD_MySql "crud_mysql"
	"crud_mysql/entity"
	"fmt"
	"testing"
	"time"
)

func TestPostRepositoryImpl_Insert(t *testing.T) {
	postRepository := NewPostRepository(CRUD_MySql.GetConnection())

	ctx := context.Background()

	post := entity.Post{
		Title:     "Corona Ibarat sagu",
		Body:      "ya Gitu deh g tau pokonya kek sagu aja",
		CreatedBy: 2,
		CreatedAt: time.Now(),
	}

	insert, err := postRepository.Insert(ctx, post)

	if err != nil {
		panic(err)
	}

	fmt.Println(insert)
}

func TestPostRepositoryImpl_FindById(t *testing.T) {
	postRepository := NewPostRepository(CRUD_MySql.GetConnection())

	ctx := context.Background()

	result, err := postRepository.FindById(ctx, 1)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
