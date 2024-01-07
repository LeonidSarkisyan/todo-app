package repository

import (
	"TodoApp"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func newAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user TodoApp.User) (int, error) {
	var id int
	query := fmt.Sprintf(
		"INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", usersTable)
	err := r.db.QueryRow(query, user.Name, user.Username, user.Password).Scan(&id)
	err = handlerErrorCreateUser(err)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUserByUsername(username, password string) (TodoApp.User, error) {
	var user TodoApp.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)
	return user, err
}

func handlerErrorCreateUser(err error) error {
	if err != nil {
		var pqErr *pq.Error
		ok := errors.As(err, &pqErr)
		if ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				logrus.Infof("такой пользователь с username уже существует: %s", pqErr.Error())
				return errors.New("такой пользователь с username уже существует")
			case "string_data_right_truncation":
				logrus.Infof("слишком длинное название поле (более 255 символов)")
				return errors.New("слишком длинное название поле (более 255 символов)")
			default:
				logrus.Infof(
					"случилась другая ошибка типа pq.Error: %s, c ошибкой: %s", pqErr.Code.Name(), pqErr.Error())
				return errors.New("ошибка базы данных")
			}
		} else {
			logrus.Errorf("ошибка в БД не типа pq.Error: %s", pqErr.Error())
			return err
		}
	}
	return nil
}
