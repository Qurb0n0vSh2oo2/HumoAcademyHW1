package repository

import (
	"bankCLI/pkg/models"
	"context"
)

func (r *Repository) AddTransfer(sender, recipient models.Account, amount float64)(err error) {
	_, err = r.Database.Exec(context.Background(), `
	INSERT 
	INTO 
		tranfer(sender_id, recipient_id, amount)
	VALUES
		($1, $2, $3)`,
	sender, recipient, amount)

	if err != nil{
		return err
	}
	return
}
