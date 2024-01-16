package transport

import "bankCLI/pkg/service"

type Transport struct {
	Svc service.ServiceInterface
}

type TransferInterface interface {
	CreateBankAccount()
	TopUpBankAccount()
	TakeOffMoney()
	ShowAccountBalance()
	TransferMoney()
}

func NewTransport(svc service.ServiceInterface) TransferInterface {
	return &Transport{
		Svc: svc,
	}
}
