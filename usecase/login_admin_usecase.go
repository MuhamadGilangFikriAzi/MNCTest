package usecase

import (
	"gokost.com/m/delivery/apprequest"
	"gokost.com/m/delivery/appresponse"
	"gokost.com/m/repository"
)

type LoginAdminUsecase interface {
	LoginAdmin(LoginData apprequest.AdminRequest) (appresponse.LoginResponse, bool, error)
}

type loginAdminUsecase struct {
	repo repository.AdminRepo
}

func (l *loginAdminUsecase) LoginAdmin(LoginData apprequest.AdminRequest) (appresponse.LoginResponse, bool, error) {
	return l.repo.Login(LoginData)
}

func NewLoginAdminUsecase(repo repository.AdminRepo) LoginAdminUsecase {
	return &loginAdminUsecase{
		repo,
	}
}
