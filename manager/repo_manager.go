package manager

import (
	"github.com/jmoiron/sqlx"
	"gokost.com/m/repository"
	"gorm.io/gorm"
)

type RepoManager interface {
	AdminRepo() repository.AdminRepo
}

type repoManager struct {
	SqlxDb *sqlx.DB
	GormDb *gorm.DB
}

func (r *repoManager) AdminRepo() repository.AdminRepo {
	return repository.NewAdminRepo(r.SqlxDb)
}

func NewRepoManager(sqlxDb *sqlx.DB, gormDb *gorm.DB) RepoManager {
	return &repoManager{
		sqlxDb,
		gormDb,
	}
}
