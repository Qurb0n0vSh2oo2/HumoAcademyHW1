package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) GetSenderBalance(name string) (balance float64, err error) {
	row := r.Database.QueryRow(context.Background(), `
	SELECT 
		balance
	FROM account
	WHERE name = $1
	`, name)

	err = row.Scan(&balance)
	if err != nil {
		if err == pgx.ErrNoRows {
			fmt.Println("нету баланса!")
		}
	}
	return
}
