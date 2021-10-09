package repository

import (
	"context"
	"crud_mysql/entity"
	"database/sql"
	"errors"
	"strconv"
)

type (
	commentRepositoryImpl struct {
		db *sql.DB
	}
)

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{db: db}
}

func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	query := "INSERT INTO  comments(user_id, post_id, comment, created_at) VALUES (?,?,?,?)"

	result, err := repository.db.ExecContext(ctx, query, comment.UserId, comment.PostId, comment.Comment, comment.CreatedAt)

	if err != nil {
		return comment, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return comment, err
	}
	comment.Id = int32(id)

	return comment, nil
}

func (repository *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	query := "SELECT comments.id,users.name, comments.comment, posts.title, comments.created_at FROM comments INNER JOIN users ON users.id = comments.user_id INNER JOIN posts ON posts.id = comments.post_id WHERE comments.id = ? LIMIT 1"

	result, err := repository.db.QueryContext(ctx, query, id)
	comment := entity.Comment{}
	post := entity.Post{}
	user := entity.User{}
	if err != nil {
		return comment, err
	}

	defer result.Close()

	if result.Next() {
		comment.PostTitle = post.Title
		comment.UserName = user.Name
		result.Scan(&comment.Id, &comment.UserName, &comment.Comment, &comment.PostTitle, &comment.CreatedAt)

		return comment, err
	} else {
		return comment, errors.New("Nilai id dari " + strconv.Itoa(int(id)) + "Tidak Ditemukan")
	}
}

func (repository *commentRepositoryImpl) FindByPost(ctx context.Context, postId int32) ([]entity.Comment, error) {
	query := "SELECT comments.id,users.name, comments.comment, posts.title, comments.created_at FROM comments INNER JOIN users ON users.id = comments.user_id INNER JOIN posts ON posts.id = comments.post_id WHERE comments.post_id = ?"

	result, err := repository.db.QueryContext(ctx, query, postId)
	comments := []entity.Comment{}
	post := entity.Post{}
	user := entity.User{}
	if err != nil {
		return comments, err
	}

	defer result.Close()

	for result.Next() {
		comment := entity.Comment{}
		comment.PostTitle = post.Title
		comment.UserName = user.Name
		result.Scan(&comment.Id, &comment.UserName, &comment.Comment, &comment.PostTitle, &comment.CreatedAt)
		comments = append(comments, comment)
	}

	return comments, nil
}

func (repository *commentRepositoryImpl) FindByUser(ctx context.Context, userId int32) ([]entity.Comment, error) {
	query := "SELECT comments.id,users.name, comments.comment, posts.title, comments.created_at FROM comments INNER JOIN users ON users.id = comments.user_id INNER JOIN posts ON posts.id = comments.post_id WHERE comments.user_id = ?"

	result, err := repository.db.QueryContext(ctx, query, userId)
	comments := []entity.Comment{}
	post := entity.Post{}
	user := entity.User{}
	if err != nil {
		return comments, err
	}

	defer result.Close()

	for result.Next() {
		comment := entity.Comment{}
		comment.PostTitle = post.Title
		comment.UserName = user.Name
		result.Scan(&comment.Id, &comment.UserName, &comment.Comment, &comment.PostTitle, &comment.CreatedAt)
		comments = append(comments, comment)
	}

	return comments, nil
}

func (repository *commentRepositoryImpl) UpdateComment(ctx context.Context, comment entity.Comment, id int32) (entity.Comment, error) {
	query := "UPDATE comments SET comment = ? WHERE id = ?"

	result, err := repository.db.QueryContext(ctx, query, comment.Comment, id)

	defer result.Close()
	if err != nil {
		return comment, err
	}

	return comment, nil
}

func (repository *commentRepositoryImpl) DeleteComment(ctx context.Context, id int32) error {
	query := "DELETE FROM comments WHERE id = ?"
	deleteData, err := repository.db.QueryContext(ctx, query, id)

	if err != nil {
		return err
	}

	defer deleteData.Close()

	return nil
}
