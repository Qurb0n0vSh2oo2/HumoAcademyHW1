package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (repo *Repository) GetCodeByUserID(userid int) (code string, err error) {

	row := repo.db.QueryRow(
		context.Background(),
		`SELECT 
			code
		FROM user_reset_password
		WHERE 
			user_id = $1 
		AND 
			is_active = true
		AND 
			expired_at > now()`, userid)

	err = row.Scan(&code)

	if err != nil {
		if err == pgx.ErrNoRows {
			return code, errors.New("Извините, но вы не оставляли запрос")
		}
	}

	return
}

/*
	SELECT
		code
	FROM user_reset_password
	WHERE
		user_id = 1
	AND
		is_active = true
	AND
		expired_at > now()

*/
