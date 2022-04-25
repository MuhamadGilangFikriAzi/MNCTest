package manager

import (
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
	"mnctest.com/api/repository"
)

type RepoManager interface {
	CustomerRepo() repository.CustomerRepo
}

type repoManager struct {
	SqlxDb *sqlx.DB
	GormDb *gorm.DB
}

func (r *repoManager) CustomerRepo() repository.CustomerRepo {
	return repository.NewCustomerRepo(r.SqlxDb)
}

func NewRepoManager(sqlxDb *sqlx.DB, gormDb *gorm.DB) RepoManager {
	return &repoManager{
		sqlxDb,
		gormDb,
	}
}
