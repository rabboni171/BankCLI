package repository

import (
	"bankCLI/pkg/errors"
	"bankCLI/pkg/models"
	"context"

	"github.com/jackc/pgx/v5"
)

func (repo *Repository) GetAccountByPhoneNumber(phoneNumber string) (account models.Account, err error) {
	row := repo.Conn.QueryRow(context.Background(),
		`SELECT id, 
       			full_name, 
      			 phone_number, 
     			  address, balance 
		FROM account 
		WHERE 
		phone_number = $1
		`, phoneNumber)

	err = row.Scan(
		&account.Id,
		&account.FullName,
		&account.PhoneNumber,
		&account.Address,
		&account.Balance,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return account, errors.ErrDataNotFound
		}
	}

	return
}
