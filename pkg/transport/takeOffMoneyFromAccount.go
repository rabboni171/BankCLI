package transport

import (
	"bankCLI/pkg/models"
	"fmt"
)

func (t *Transport) TakeOffMoney() {
	var account models.Account
	var amount float64

	fmt.Println("Введите номер телефона:")
	fmt.Scan(&account.PhoneNumber)

	fmt.Println("Какую сумму хотите снять:")
	fmt.Scan(&amount)

	err := t.Svc.TakeOffMoneyFromAccount(account.PhoneNumber, amount)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Уcпешно снято!")
}
