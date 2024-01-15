package repository

import (
  "bankCLI/pkg/models"
  "context"
  "errors"
  "github.com/jackc/pgx/v5"
)

func (repo *Repository) GetAccount(name string) (models.Account, error) {
  account := models.Account{
    City: &models.City{},
  }
  row := repo.Database.QueryRow(context.Background(), `
      SELECT id, name, balance, phone_number, city_id 
        FROM account WHERE name = $1`, name)

  err := row.Scan(&account.Id, &account.Name, &account.Balance, &account.PhoneNumber, &account.City.Id)

  if err != nil {
    if errors.Is(err, pgx.ErrNoRows) {
      return models.Account{}, errors.New("not found")
    }
    return models.Account{}, err
  }

  return account, nil
}
