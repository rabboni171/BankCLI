package repository

import (
	"bankCLI/pkg/errors"
	"context"
)

func (repo *Repository) GetAccountBalance(phoneNumber string) (balance float64, err error) {
	row := repo.Conn.QueryRow(context.Background(), `
	SELECT 
		balance
	FROM 
		account
	WHERE 
		phone_number = $1
	`, phoneNumber)

	err = row.Scan(&balance)

	if err != nil {
		return 0, errors.ErrNoAccount
	}

	return

}
