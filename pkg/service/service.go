package service

import (
	"bankCLI/pkg/models"
	"bankCLI/pkg/repository"
)

type Service struct {
	Repo repository.RepositoryInterface
}

type ServiceInterface interface {
	CreateBankAccount(account models.Account) error
	TopUpBankAccount(phoneNumber string, amount float64) (err error)
	TakeOffMoneyFromAccount(phoneNumber string, amount float64) (err error)
	ShowAccountBalance(account models.Account) (name string, balance float64, err error)
	TransferMoney(senderPhone, recipientPhone string, amount float64) error
}

func NewService(repo repository.RepositoryInterface) ServiceInterface {
	return &Service{
		Repo: repo,
	}
}
