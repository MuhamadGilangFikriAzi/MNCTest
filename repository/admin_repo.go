package repository

import (
	"github.com/jmoiron/sqlx"
	"gokost.com/m/delivery/apprequest"
	"gokost.com/m/delivery/appresponse"
	"gokost.com/m/utility"
)

type AdminRepo interface {
	Login(loginData apprequest.AdminRequest) (appresponse.LoginResponse, bool, error)
}

type adminRepo struct {
	sqlDb *sqlx.DB
}

func (a *adminRepo) Login(loginData apprequest.AdminRequest) (appresponse.LoginResponse, bool, error) {
	var dataPassword appresponse.LoginResponse
	err := a.sqlDb.Get(&dataPassword, "select name, username, password from admin_kost where username = $1 and is_active = true", loginData.Username)
	is_available := utility.CheckPasswordHash(loginData.Password, dataPassword.Password)
	if err != nil {
		return dataPassword, is_available, err
	}
	return dataPassword, is_available, nil
}

func NewAdminRepo(sqlDb *sqlx.DB) AdminRepo {
	return &adminRepo{
		sqlDb,
	}
}
