package service

import "bankCLI/pkg/errors"

func (s *Service) TopUpBankAccount(phoneNumber string, amount float64) (err error) {
	_, err = s.Repo.GetAccountByPhoneNumber(phoneNumber)
	if err != nil {
		return errors.ErrNoAccount
	}
	err = s.Repo.TopUpBankAccount(phoneNumber, amount)
	if err != nil {
		return err
	}

	return nil
}
