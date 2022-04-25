package manager

import "mnctest.com/api/usecase"

type UseCaseManager interface {
	LoginCustomerUseCase() usecase.LoginCustomerUsecase
	CustomerUseCase() usecase.CustomerUseCase
	TransactionUseCase() usecase.BalanceTransferUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (u *useCaseManager) LoginCustomerUseCase() usecase.LoginCustomerUsecase {
	return usecase.NewLoginCustomerUsecase(u.repo.CustomerRepo())
}

func (u *useCaseManager) CustomerUseCase() usecase.CustomerUseCase {
	return usecase.NewCustomerUsecase(u.repo.CustomerRepo())
}

func (u *useCaseManager) TransactionUseCase() usecase.BalanceTransferUseCase {
	return usecase.NewBlanceTransferUseCase(u.repo.TransactionRepo())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{
		repo,
	}
}
