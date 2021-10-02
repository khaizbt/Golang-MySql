package repository

import (
	"context"
	CRUD_MySql "crud_mysql"
	"crud_mysql/entity"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"time"
)

func TestUserInsert(t *testing.T) {
	userRepository := NewUserRepository(CRUD_MySql.GetConnection())

	ctx := context.Background()
	user := entity.User{
		Name:      "Kukuruyuk",
		Email:     "tamamizadada@gmail,com",
		CreatedAt: time.Now(),
	}

	insert, err := userRepository.Insert(ctx, user)

	if err != nil {
		panic(err)
	}

	fmt.Println(insert)

}

func TestUserRepositoryImpl_FindById(t *testing.T) {
	userRepository := NewUserRepository(CRUD_MySql.GetConnection())

	ctx := context.Background()

	result, err := userRepository.FindById(ctx, 2)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestUserRepositoryImpl_FindAll(t *testing.T) {
	userRepository := NewUserRepository(CRUD_MySql.GetConnection())

	ctx := context.Background()

	results, err := userRepository.FindAll(ctx)

	if err != nil {
		panic(err)
	}

	for _, result := range results {
		fmt.Println(result)
	}
}

func TestUserRepositoryImpl_DeleteById(t *testing.T) {
	userRepository := NewUserRepository(CRUD_MySql.GetConnection())

	ctx := context.Background()

	err := userRepository.DeleteById(ctx, 1)

	if err != nil {
		panic(err)
	}

	fmt.Println("Sukses Delete Data")
}

func TestUserRepositoryImpl_UpdateData(t *testing.T) {
	userRepository := NewUserRepository(CRUD_MySql.GetConnection())

	ctx := context.Background()

	user := entity.User{
		Name:  "Halo Bos",
		Email: "Siap",
	}

	update, err := userRepository.UpdateData(ctx, 3, user)

	if err != nil {
		panic(err)
	}

	fmt.Println(update)
}
