package repository

import (
	"bankCLI/pkg/config"
	"bankCLI/pkg/models"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	Database *pgx.Conn
	Percent config.Config
}

type RepositoryInterface interface {
	AddAccount(name string, balance float64, phone_number, password string, city models.City) (err error)
	AddCity(name, region string) (err error)
	AddTransfer(sender, recipient models.Account, amount float64)(err error)
	// ChangeAccountsBalance(account *models.Account, newBalance float64)
	GetAccount(name string) (account models.Account, err error)
	GetCity(name string) (city models.City, err error)
	GetPercent() float64
	TopUpProfitAccount(amount float64)
	GetSenderBalance(name string) (balance float64, err error)
	ChangeSenderBalance(name string, amount float64) error
	ChangeRecipientBalance(senderName string, amount float64) error
}

func NewRepository(db *pgx.Conn, percent config.Config) RepositoryInterface {
	return &Repository{
		Database: db,
		Percent: percent,
	}
}
