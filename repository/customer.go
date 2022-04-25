package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"mnctest.com/api/delivery/apprequest"
	"mnctest.com/api/delivery/appresponse"
	"mnctest.com/api/model"
	"mnctest.com/api/utility"
	"time"
)

type CustomerRepo interface {
	Login(loginData apprequest.CustomerRequest) (appresponse.CustomerResponse, bool, error)
	GetListCustomer(dataMeta apprequest.Meta) ([]appresponse.CustomerResponse, apprequest.Meta, error)
	GetCustomerById(getId string) (appresponse.CustomerResponse, error)
	CreatedCustomer(data apprequest.CustomerRequest) (appresponse.CustomerResponse, error)
	UpdateCustomer(updateId string, dataUpdate apprequest.CustomerRequest) error
	DeleteCustomer(deleteId string) error
	RegisterCustomer(updateId string) error
}

type customerRepo struct {
	db *sqlx.DB
}

func (a *customerRepo) Login(loginData apprequest.CustomerRequest) (appresponse.CustomerResponse, bool, error) {
	var dataPassword model.Customer
	var dataReturn appresponse.CustomerResponse
	err := a.db.Get(&dataPassword, "select name, username, password from customer where username = $1 and is_register = true limit 1", loginData.Username)
	fmt.Println(dataPassword)
	is_available := utility.CheckPasswordHash(loginData.Password, dataPassword.Password)
	dataReturn.Name = dataPassword.Name
	dataReturn.Username = dataPassword.Username
	if err != nil {
		return dataReturn, is_available, err
	}
	return dataReturn, is_available, nil
}

func (repo *customerRepo) GetListCustomer(dataMeta apprequest.Meta) ([]appresponse.CustomerResponse, apprequest.Meta, error) {
	var data []model.Customer
	var dataReturn []appresponse.CustomerResponse
	err := repo.db.Select(&data, "select id, name, type_customer_id, balance from customer where deleted_at is null and is_register = true limit $1 offset $2", dataMeta.Limit, dataMeta.Skip)
	if err != nil {
		return nil, dataMeta, err
	}
	var count int
	errCount := repo.db.Get(&count, "select count(*) from customer where is_register = true and deleted_at is null")
	if errCount != nil {
		return nil, dataMeta, errCount
	}
	for _, customer := range data {
		dataReturn = append(dataReturn, repo.setCustomerResp(customer))
	}
	dataMeta.Total = count
	return dataReturn, dataMeta, nil
}

func (repo *customerRepo) setCustomerResp(data model.Customer) appresponse.CustomerResponse {
	var customerResp appresponse.CustomerResponse
	var typeCustomer appresponse.TypeCustomerResponse
	repo.db.Get(&typeCustomer, "select name from type_customer where id = $1", data.TypeCustomerID)
	customerResp.CustomerId = data.Id
	customerResp.Name = data.Name
	customerResp.Balance = data.Balance
	customerResp.BalanceFormated = utility.CunrrencyFormat("Rp.", data.Balance)
	customerResp.TypeCustomer = typeCustomer
	return customerResp
}

func (repo *customerRepo) GetCustomerById(getId string) (appresponse.CustomerResponse, error) {
	var data model.Customer
	err := repo.db.Get(&data, "select id, name, type_customer_id, balance from customer where id = $1 and is_register = true and deleted_at is null", getId)
	if err != nil {
		return appresponse.CustomerResponse{}, err
	}
	return repo.setCustomerResp(data), nil
}

func (repo *customerRepo) CreatedCustomer(data apprequest.CustomerRequest) (appresponse.CustomerResponse, error) {
	uuid := utility.GenerateUUID()
	password, errPass := utility.HashPassword(data.Password)
	if errPass != nil {
		return appresponse.CustomerResponse{}, errPass
	}
	tx := repo.db.MustBegin()

	tx.MustExec("insert into customer(id, name, balance, type_customer_id, username, password) values($1, $2, $3, $4, $5, $6)", uuid, data.Name, 0, data.TypeCustomerId, data.Username, password)
	err := tx.Commit()
	if err != nil {
		//logger.SendLogToDiscord("err query add", err)
		fmt.Println(err.Error())
		return appresponse.CustomerResponse{}, err
	}
	var dataCreate model.Customer
	errGet := repo.db.Get(&dataCreate, "select id, name, type_customer_id, balance from customer where id = $1", uuid)
	dataReturn := repo.setCustomerResp(dataCreate)
	if errGet != nil {
		//logger.SendLogToDiscord("error get data create", errGet)
		return dataReturn, errGet
	}
	return dataReturn, nil
}

func (repo *customerRepo) UpdateCustomer(updateId string, dataUpdate apprequest.CustomerRequest) error {
	_, err := repo.db.Query("update customer set name = $1, type_customer_id = $2 where id = $3", dataUpdate.Name, dataUpdate.TypeCustomerId, updateId)
	if err != nil {
		return err
	}
	return nil
}

func (repo *customerRepo) RegisterCustomer(updateId string) error {
	fmt.Println(updateId)
	_, err := repo.db.Query("update customer set is_register = true where id = $1", updateId)
	if err != nil {
		return err
	}
	return nil
}

func (repo *customerRepo) DeleteCustomer(deleteId string) error {
	thisTime := time.Now()
	_, err := repo.db.Query("UPDATE customer SET deleted_at = $1 WHERE id = $2", thisTime, deleteId)
	if err != nil {
		return err
	}
	return nil
}

func NewCustomerRepo(sqlDb *sqlx.DB) CustomerRepo {
	return &customerRepo{
		sqlDb,
	}
}
