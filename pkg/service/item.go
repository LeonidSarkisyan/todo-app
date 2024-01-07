package service

import (
	"TodoApp"
	"TodoApp/pkg/repository"
)

type ItemService struct {
	repository     *repository.Repository
	listRepository *repository.TodoList
}

func NewItemService(repo *repository.Repository) *ItemService {
	return &ItemService{repository: repo}
}

func (s *ItemService) Create(userId, listId int, item TodoApp.TodoItem) (int, error) {
	_, err := s.repository.TodoList.GetById(userId, listId)
	if err != nil {
		return 0, err
	}
	return s.repository.TodoItem.Create(listId, item)
}

func (s *ItemService) GetAll(userId, listId int) ([]TodoApp.TodoItem, error) {
	_, err := s.repository.TodoList.GetById(userId, listId)
	if err != nil {
		return nil, err
	}
	return s.repository.TodoItem.GetAll(listId)
}

func (s *ItemService) GetById(userId, listId, itemId int) (TodoApp.TodoItem, error) {
	_, err := s.repository.TodoList.GetById(userId, listId)
	if err != nil {
		return TodoApp.TodoItem{}, err
	}
	_, err = s.repository.TodoItem.GetById(itemId, listId)
	return s.repository.TodoItem.GetById(itemId, listId)
}

func (s *ItemService) Update(userId, listId, itemId int, input TodoApp.UpdateItemInput) error {
	err := input.Validate()
	if err != nil {
		return err
	}
	_, err = s.repository.TodoList.GetById(userId, listId)
	if err != nil {
		return err
	}
	return s.repository.TodoItem.Update(userId, itemId, input)
}

func (s *ItemService) Delete(userId, listId, itemId int) error {
	_, err := s.repository.TodoList.GetById(userId, listId)
	if err != nil {
		return err
	}
	return s.repository.TodoItem.Delete(userId, itemId)
}
