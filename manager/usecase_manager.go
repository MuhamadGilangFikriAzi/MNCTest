package manager

import "gokost.com/m/usecase"

type UseCaseManager interface {
	LoginAdminUseCase() usecase.LoginAdminUsecase
}

type useCaseManager struct {
	repo RepoManager
}

func (u *useCaseManager) LoginAdminUseCase() usecase.LoginAdminUsecase {
	return usecase.NewLoginAdminUsecase(u.repo.AdminRepo())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{
		repo,
	}
}
