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

func (repository *postRepositoryImpl) FindAll(ctx context.Context) ([]entity.Post, error) {
	query := "SELECT posts.id, posts.title, posts.body, posts.created_at, users.name FROM posts INNER JOIN users ON users.id = posts.created_by"

	result, err := repository.DB.QueryContext(ctx, query)
	user := entity.User{}
	posts := []entity.Post{}
	if err != nil {
		return posts, err
	}

	defer result.Close()

	for result.Next() {
		post := entity.Post{}
		post.UserName = user.Name
		result.Scan(&post.Id, &post.Title, &post.Body, &post.CreatedAt, &post.UserName)

		posts = append(posts, post)
	}

	return posts, nil
}

func (repository *postRepositoryImpl) UpdatePost(ctx context.Context, post entity.Post, id int32) (entity.Post, error) {
	query := "UPDATE posts SET title = ?, body = ? WHERE id = ?"

	result, err := repository.DB.QueryContext(ctx, query, post.Title, post.Body, id)

	defer result.Close()
	if err != nil {
		return post, err
	}

	return post, nil
}

func (repository *postRepositoryImpl) FindPostByUser(ctx context.Context, userId int32) ([]entity.Post, error) {
	query := "SELECT posts.id, posts.title, posts.body, posts.created_at, users.name FROM posts INNER JOIN users ON users.id = posts.created_by WHERE created_by = ?"

	result, err := repository.DB.QueryContext(ctx, query, userId)

	defer result.Close()
	posts := []entity.Post{}
	user := entity.User{}
	if err != nil {
		return posts, err
	}

	for result.Next() {
		post := entity.Post{}
		post.UserName = user.Name
		result.Scan(&post.Id, &post.Title, &post.Body, &post.CreatedAt, &post.UserName)

		posts = append(posts, post)
	}

	return posts, err
}

func (repository *postRepositoryImpl) DeletePost(ctx context.Context, id int32) error {
	query := "DELETE FROM posts WHERE id = ?"
	deleteData, err := repository.DB.QueryContext(ctx, query, id)

	if err != nil {
		return err
	}

	defer deleteData.Close()

	return nil
}
