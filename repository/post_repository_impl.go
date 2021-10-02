package repository

import (
	"context"
	"crud_mysql/entity"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
)

type postRepositoryImpl struct { //agar waktu query tidak memanggil db berulang kali
	DB *sql.DB
}

func NewPostRepository(db *sql.DB) PostRepository {
	return &postRepositoryImpl{DB: db}
}

func (repository *postRepositoryImpl) Insert(ctx context.Context, post entity.Post) (entity.Post, error) {
	query := "INSERT INTO posts(title, body, created_by, created_at) VALUES (?,?,?,?)"
	result, err := repository.DB.ExecContext(ctx, query, post.Title, post.Body, post.CreatedBy, post.CreatedAt)

	if err != nil {
		return post, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return post, err
	}
	post.Id = int32(id)
	return post, nil
}

func (repository *postRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Post, error) {
	query := "SELECT posts.id, posts.title, posts.body, posts.created_at, users.name FROM posts INNER JOIN users ON users.id = posts.created_by WHERE posts.id = ? LIMIT 1"

	result, err := repository.DB.QueryContext(ctx, query, id)
	user := entity.User{}
	post := entity.Post{}

	if err != nil {
		return post, err
	}

	defer result.Close()

	if result.Next() {
		fmt.Println(post.UserName)
		post.UserName = user.Name
		result.Scan(&post.Id, &post.Title, &post.Body, &post.CreatedAt, &post.UserName)

		return post, nil
	} else {
		return post, errors.New("Id " + strconv.Itoa(int(id)) + "Tidak Ditemukan")
	}
}
