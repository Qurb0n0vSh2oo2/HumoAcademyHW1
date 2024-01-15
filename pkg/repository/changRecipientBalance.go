package repository

import "context"

func (r *Repository) ChangeRecipientBalance(senderName string, amount float64) error {
	_, err := r.Database.Exec(context.Background(), `
	UPDATE account
	SET balance = balance + $1
	WHERE name = $2
	`, amount, senderName)

	return err
}
