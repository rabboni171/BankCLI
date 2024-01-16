package transport

import (
	"bankCLI/pkg/models"
	"fmt"
)

func (t *Transport) TopUpBankAccount() {
	var account models.Account
	var amount float64

	fmt.Println("Введите номер телефона:")
	fmt.Scan(&account.PhoneNumber)

	fmt.Println("Введите сумму для пополнения счёта:")
	fmt.Scan(&amount)

	err := t.Svc.TopUpBankAccount(account.PhoneNumber, amount)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Транзакция успешно завершена!")
}
