package main

import (
	"bankCLI/pkg/config"
	"bankCLI/pkg/database"
	"bankCLI/pkg/repository"
	"bankCLI/pkg/service"
	"bankCLI/pkg/transport"
	"fmt"
)

func main() {
	conf := config.NewConfig()
	db := database.NewDatabase(conf)
	repo := repository.NewRepository(db, *conf)
	svc := service.NewService(repo)
	transp := transport.NewTransport(svc)

	for {
		var choice int

		fmt.Println("<<BANK CLI WITH SQL>>")
		fmt.Println("1. Создать счет в банке")
		fmt.Println("2. Пополнить счёт")
		fmt.Println("3. Снять деньги")
		fmt.Println("4. Показать счёт")
		fmt.Println("5. Перевод денег")
		fmt.Println("6. Выйти")

		fmt.Scan(&choice)

		if choice == 1 {
			transp.CreateBankAccount()
		} else if choice == 2 {
			transp.TopUpBankAccount()
		} else if choice == 3 {
			transp.TakeOffMoney()
		} else if choice == 4 {
			transp.ShowAccountBalance()
		} else if choice == 5 {
			transp.TransferMoney()
		} else if choice == 6 {
			return
		}
	}
}
