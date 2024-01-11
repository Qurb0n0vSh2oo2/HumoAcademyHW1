package repository

import "context"

func (repo *Repository) ForgotPassword(user_id int, code string) (err error) {
	_, err = repo.db.Exec(
		context.Background(),
		"INSERT INTO user_reset_password(code, user_id, expired_at) VALUES ($1, $2, now() + interval '1' day)",
		code,
		user_id,
	)

	return
}
