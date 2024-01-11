package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) Login(username, password string) (id int, err error) {
	var db_password string

	row := r.db.QueryRow(context.Background(),
		"SELECT password, id FROM users WHERE username = $1",
		username)

	err = row.Scan(&db_password, &id)

	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, errors.New("пользователь не найден")
		}

		return 0, err
	}

	if db_password != password {
		return 0, errors.New("неправильный пароль")
	}

	return id, nil
}

// db.Query // получить несколько записей из бд
// db.QueryRow // получить одну запись из бд
//db.Exec // отправить запрос но не ожидать ничего из бд
