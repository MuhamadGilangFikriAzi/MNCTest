package usecase

import (
	"mnctest.com/api/delivery/apprequest"
	"mnctest.com/api/delivery/appresponse"
	"mnctest.com/api/repository"
)

type LoginCustomerUsecase interface {
	LoginCustomer(LoginData apprequest.CustomerRequest) (appresponse.CustomerResponse, bool, error)
}

type loginCustomerUsecase struct {
	repo repository.CustomerRepo
}

func (l *loginCustomerUsecase) LoginCustomer(LoginData apprequest.CustomerRequest) (appresponse.CustomerResponse, bool, error) {
	return l.repo.Login(LoginData)
}

func NewLoginCustomerUsecase(repo repository.CustomerRepo) LoginCustomerUsecase {
	return &loginCustomerUsecase{
		repo,
	}
}
