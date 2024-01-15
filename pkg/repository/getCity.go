package repository

import (
	"bankCLI/pkg/models"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (repo *Repository) GetCity(name string) (models.City, error) {
	city := models.City{}
	row := repo.Database.QueryRow(context.Background(), `
    SELECT 
      id, name, region
    FROM city
    WHERE name = $1
  `, name)

	err := row.Scan(&city.Id, &city.Name, &city.Region)

	if err != nil {
		if err == pgx.ErrNoRows {
			return models.City{}, errors.New("not found")
		}
	}

	return city, nil
}
