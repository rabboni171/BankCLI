package repository

import (
	"bankCLI/pkg/config"
	"bankCLI/pkg/models"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	Conn    *pgx.Conn
	Percent config.Config
}

type RepositoryInterface interface {
	CreateBankAccount(account models.Account) (err error)
	GetAccountByName(name string) (account models.Account, err error)
	TopUpBankAccount(phoneNumber string, amount float64) (err error)
	GetAccountByPhoneNumber(phoneNumber string) (account models.Account, err error)
	GetAccountBalance(phoneNumber string) (balance float64, err error)
	TakeOffMoneyFromAccount(phoneNumber string, amount float64) (err error)
	GetPercent() float64
	ChangeAccountBalance(account models.Account) (err error)
	AddProfitAccount(amount float64) error
	AddTransfer(sender *models.Account, recipient *models.Account, amount float64) (err error)
}

func NewRepository(conn *pgx.Conn, percent config.Config) RepositoryInterface {
	return &Repository{
		Conn:    conn,
		Percent: percent,
	}
}
