package repository

import "context"

func (repo *Repository) DeactivateByCode(reset_asctivnes string) (err error) {
	_, err = repo.db.Exec(context.Background(), `
		UPDATE user_reset_password
		SET is_active = false
		WHERE code = $1
	`, reset_asctivnes)

	return
}
