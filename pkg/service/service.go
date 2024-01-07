package service

import (
	"TodoApp"
	"TodoApp/pkg/repository"
)

type Authorization interface {
	CreateUser(user TodoApp.User) (int, error)
	LoginUser(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list TodoApp.TodoList) (int, error)
	GetAll(userId int) ([]TodoApp.TodoList, error)
	GetById(userId, listId int) (TodoApp.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input TodoApp.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, item TodoApp.TodoItem) (int, error)
	GetAll(userId, listId int) ([]TodoApp.TodoItem, error)
	GetById(userId, listId, itemId int) (TodoApp.TodoItem, error)
	Update(userId, listId, itemId int, input TodoApp.UpdateItemInput) error
	Delete(userId, listId, itemId int) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo),
		TodoList:      NewListService(repo),
		TodoItem:      NewItemService(repo),
	}
}
