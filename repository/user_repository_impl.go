package repository

import (
	"context"
	"crud_mysql/entity"
	"database/sql"
	"errors"
	"strconv"
)

type userRepositoryImpl struct { //agar waktu query tidak memanggil db berulang kali
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepositoryImpl{DB: db}
}

func (repository *userRepositoryImpl) Insert(ctx context.Context, user entity.User) (entity.User, error) {
	script := "INSERT INTO users(name, email, created_at) VALUES (?, ?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, user.Name, user.Email, user.CreatedAt)
	if err != nil {
		return user, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return user, err
	}
	user.Id = int32(id)
	return user, nil
}

func (repository *userRepositoryImpl) FindById(ctx context.Context, id int32) (entity.User, error) {
	query := "SELECT id, name, email, created_at FROM users WHERE id = ? LIMIT 1"
	result, err := repository.DB.QueryContext(ctx, query, id)
	user := entity.User{}
	if err != nil {
		return user, err
	}

	defer result.Close()

	if result.Next() {
		result.Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt)
		return user, nil
	} else {
		return user, errors.New("Id " + strconv.Itoa(int(id)) + "Tidak Ditemukan")
	}

}

func (repository *userRepositoryImpl) FindAll(ctx context.Context) ([]entity.User, error) {
	query := "SELECT id, name, email, created_at FROM users"
	result, err := repository.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer result.Close()

	users := []entity.User{}

	for result.Next() {
		user := entity.User{}
		result.Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt)
		users = append(users, user)
	}

	return users, nil
}

func (repository *userRepositoryImpl) DeleteById(ctx context.Context, id int32) error {
	query := "DELETE FROM users WHERE id = ?"
	delete, err := repository.DB.QueryContext(ctx, query, id)

	if err != nil {
		return err
	}

	defer delete.Close()

	return nil
}

func (repository *userRepositoryImpl) UpdateData(ctx context.Context, id int32, user entity.User) (entity.User, error) {
	query := "UPDATE users SET name = ?, email = ? WHERE id = ?"

	result, err := repository.DB.QueryContext(ctx, query, user.Name, user.Email, id)
	defer result.Close()
	if err != err {
		return user, err
	}

	return user, nil
}
