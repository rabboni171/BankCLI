package transport

import (
	"bankCLI/pkg/models"
	"fmt"
)

func (t *Transport) ShowAccountBalance() {
	var account models.Account

	fmt.Println("Введите номер телефона")

	fmt.Scan(&account.PhoneNumber)

	name, balance, err := t.Svc.ShowAccountBalance(account)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Баланс счета %v : %v \n", name, balance)
}
