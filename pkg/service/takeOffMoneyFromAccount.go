package service

import (
	"bankCLI/pkg/errors"
	"fmt"
)

func (s *Service) TakeOffMoneyFromAccount(phoneNumber string, amount float64) (err error) {
	_, err = s.Repo.GetAccountByPhoneNumber(phoneNumber)
	if err != nil {
		return errors.ErrNoAccount
	}

	balance, err := s.Repo.GetAccountBalance(phoneNumber)
	if balance < amount {
		return errors.ErrNotEnoughMoney
	}
	err = s.Repo.TakeOffMoneyFromAccount(phoneNumber, amount)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
