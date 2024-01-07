package repository

import (
	"TodoApp"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

type ItemPostgres struct {
	db *sqlx.DB
}

func newItemPostgres(db *sqlx.DB) *ItemPostgres {
	return &ItemPostgres{db: db}
}

func (r *ItemPostgres) Create(listId int, item TodoApp.TodoItem) (int, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return 0, err
	}

	defer tx.Rollback()

	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id",
		todoItemsTable)

	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	err = row.Scan(&itemId)
	if err != nil {
		return 0, err
	}

	createListItemsQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) VALUES ($1, $2)",
		listsItemsTable)

	_, err = tx.Exec(createListItemsQuery, listId, itemId)
	if err != nil {
		return 0, err
	}

	return itemId, tx.Commit()
}

func (r *ItemPostgres) GetAll(listId int) ([]TodoApp.TodoItem, error) {
	var items []TodoApp.TodoItem

	query := fmt.Sprintf("SELECT * FROM %s WHERE list_id = $1", todoItemsTable)

	err := r.db.Select(&items, query, listId)
	return items, err
}

func (r *ItemPostgres) GetById(itemId, listId int) (TodoApp.TodoItem, error) {
	var item TodoApp.TodoItem
	query := fmt.Sprintf("SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li ON ti.id = li.item_id WHERE li.list_id = $1 AND li.item_id = $2",
		todoItemsTable, listsItemsTable)
	err := r.db.Get(&item, query, listId, itemId)
	return item, err
}

func (r *ItemPostgres) Update(userId, itemId int, input TodoApp.UpdateItemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *input.Done)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s ti SET %s FROM %s li, %s ul
									WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $%d AND ti.id = $%d`,
		todoItemsTable, setQuery, listsItemsTable, usersListsTable, argId, argId+1)
	args = append(args, userId, itemId)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *ItemPostgres) Delete(userId, itemId int) error {
	query := fmt.Sprintf(`DELETE FROM %s ti USING %s li, %s ul 
									WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $1 AND ti.id = $2`,
		todoItemsTable, listsItemsTable, usersListsTable)
	_, err := r.db.Exec(query, userId, itemId)
	return err
}
