package repository

import (
	"bankCLI/pkg/models"
	"context"
)

func (repo *Repository) AddAccount(name string, balance float64, phoneNumber, password string, city models.City) (err error) {
	_, err = repo.Database.Exec(context.Background(), `
	  INSERT INTO account(
		name,
		balance,
		phone_number,
		  password,
		city_id
	  )
	  VALUES(
		$1,
		$2,
		$3,
		$4,
		$5
	  );
	`, name, balance, phoneNumber, password, city.Id)

	return
}
