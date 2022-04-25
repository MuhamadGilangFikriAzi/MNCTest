package usecase

import (
	"mnctest.com/api/delivery/apprequest"
	"mnctest.com/api/delivery/appresponse"
	"mnctest.com/api/repository"
)

type BalanceTransferUseCase interface {
	BalanceTransfer(data apprequest.TransactionRequest) (appresponse.TransactionResponse, string, error)
}

type balanceTransferUseCase struct {
	repo repository.Transaction
}

func (usecase *balanceTransferUseCase) BalanceTransfer(data apprequest.TransactionRequest) (appresponse.TransactionResponse, string, error) {
	return usecase.repo.BalanceTransfer(data)
}

func NewBlanceTransferUseCase(repo repository.Transaction) BalanceTransferUseCase {
	return &balanceTransferUseCase{repo}
}
