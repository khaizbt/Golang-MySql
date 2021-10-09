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
		Title:     "Corona Turun, PPKM Ngga HEHE",
		Body:      "Mana saya tau lah",
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

	result, err := postRepository.FindById(ctx, 2)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestPostRepositoryImpl_FindAll(t *testing.T) {
	postRepository := NewPostRepository(CRUD_MySql.GetConnection())

	ctx := context.Background()

	results, err := postRepository.FindAll(ctx)

	if err != nil {
		panic(err)
	}

	for _, result := range results {
		fmt.Println(result)
	}
}

func TestPostRepositoryImpl_UpdatePost(t *testing.T) {
	postRepository := NewPostRepository(CRUD_MySql.GetConnection())
	ctx := context.Background()

	post := entity.Post{
		Title: "Siap Bang",
		Body:  "Siappppppppppp",
	}

	result, err := postRepository.UpdatePost(ctx, post, 2)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestPostRepositoryImpl_FindPostByUser(t *testing.T) {
	postRepository := NewPostRepository(CRUD_MySql.GetConnection())
	ctx := context.Background()

	results, err := postRepository.FindPostByUser(ctx, 2)
	if err != nil {
		panic(err)
	}

	for _, result := range results {
		fmt.Println(result)
	}
}

func TestPostRepositoryImpl_DeletePost(t *testing.T) {
	postRespository := NewPostRepository(CRUD_MySql.GetConnection())

	ctx := context.Background()

	err := postRespository.DeletePost(ctx, 1)

	if err != nil {
		panic(err)
	}

	fmt.Println("Sukses Delete Data")
}
