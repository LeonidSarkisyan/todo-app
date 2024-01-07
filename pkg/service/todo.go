package service

import (
	"TodoApp"
	"TodoApp/pkg/repository"
)

type ListService struct {
	repository *repository.Repository
}

func NewListService(repository *repository.Repository) *ListService {
	return &ListService{repository: repository}
}

func (s *ListService) Create(userId int, list TodoApp.TodoList) (int, error) {
	return s.repository.TodoList.Create(userId, list)
}

func (s *ListService) GetAll(userId int) ([]TodoApp.TodoList, error) {
	return s.repository.TodoList.GetAll(userId)
}

func (s *ListService) GetById(userId, listId int) (TodoApp.TodoList, error) {
	return s.repository.TodoList.GetById(userId, listId)
}

func (s *ListService) Delete(userId, listId int) error {
	return s.repository.TodoList.Delete(userId, listId)
}

func (s *ListService) Update(userId, listId int, input TodoApp.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repository.TodoList.Update(userId, listId, input)
}
