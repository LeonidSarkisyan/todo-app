package TodoApp

import "errors"

type TodoList struct {
	Id          int    `json:"id"`
	Title       string `json:"title" binding:"required" db:"title"`
	Description string `json:"description" binding:"required" db:"description"`
}

type UsersList struct {
	Id         int `db:"id"`
	UserId     int `db:"user_id"`
	TodoListId int `db:"list_id"`
}

type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title" binding:"required" db:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type UpdateItemInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}

func (i UpdateItemInput) Validate() error {
	if i.Title == nil && i.Description == nil && i.Done == nil {
		return errors.New("update structure has no values")
	}
	return nil
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (i UpdateListInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
