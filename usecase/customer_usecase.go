package usecase

import (
	"mnctest.com/api/delivery/apprequest"
	"mnctest.com/api/delivery/appresponse"
	"mnctest.com/api/repository"
)

type CustomerUseCase interface {
	GetAllCustomer(dataMeta apprequest.Meta) ([]appresponse.CustomerResponse, apprequest.Meta, error)
	CreateCustomer(data apprequest.CustomerRequest) (appresponse.CustomerResponse, error)
	DetailCustomer(detailId string) (appresponse.CustomerResponse, error)
	UpdateCustomer(updateId string, dataUpdate apprequest.CustomerRequest) error
	DeleteCustomer(deleteId string) error
	RegisterCustomer(registerId string) error
}

type customerUsecase struct {
	repo repository.CustomerRepo
}

func (usecase *customerUsecase) GetAllCustomer(dataMeta apprequest.Meta) ([]appresponse.CustomerResponse, apprequest.Meta, error) {
	return usecase.repo.GetListCustomer(dataMeta)
}

func (usecase *customerUsecase) CreateCustomer(data apprequest.CustomerRequest) (appresponse.CustomerResponse, error) {
	return usecase.repo.CreatedCustomer(data)
}

func (usecase *customerUsecase) DetailCustomer(detailId string) (appresponse.CustomerResponse, error) {
	return usecase.repo.GetCustomerById(detailId)
}

func (usecase *customerUsecase) UpdateCustomer(updateId string, dataUpdate apprequest.CustomerRequest) error {
	return usecase.repo.UpdateCustomer(updateId, dataUpdate)
}

func (usecase *customerUsecase) DeleteCustomer(deleteId string) error {
	return usecase.repo.DeleteCustomer(deleteId)
}

func (usecase *customerUsecase) RegisterCustomer(registerId string) error {
	return usecase.repo.RegisterCustomer(registerId)
}

func NewCustomerUsecase(repo repository.CustomerRepo) CustomerUseCase {
	return &customerUsecase{repo}
}
