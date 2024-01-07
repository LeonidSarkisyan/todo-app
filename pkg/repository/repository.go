package repository

import (
	"TodoApp"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user TodoApp.User) (int, error)
	GetUserByUsername(username, password string) (TodoApp.User, error)
}

type TodoList interface {
	Create(userId int, list TodoApp.TodoList) (int, error)
	GetAll(userId int) ([]TodoApp.TodoList, error)
	GetById(userId, listId int) (TodoApp.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input TodoApp.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item TodoApp.TodoItem) (int, error)
	GetAll(listId int) ([]TodoApp.TodoItem, error)
	GetById(itemId, listId int) (TodoApp.TodoItem, error)
	Update(userId, itemId int, input TodoApp.UpdateItemInput) error
	Delete(userId, itemId int) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: newAuthPostgres(db),
		TodoList:      newListPostgres(db),
		TodoItem:      newItemPostgres(db),
	}
}
