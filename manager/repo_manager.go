package manager

import (
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
	"mnctest.com/api/repository"
)

type RepoManager interface {
	CustomerRepo() repository.CustomerRepo
	TransactionRepo() repository.Transaction
}

type repoManager struct {
	SqlxDb *sqlx.DB
	GormDb *gorm.DB
}

func (r *repoManager) CustomerRepo() repository.CustomerRepo {
	return repository.NewCustomerRepo(r.SqlxDb)
}

func (r *repoManager) TransactionRepo() repository.Transaction {
	return repository.NewTransaction(r.SqlxDb, repository.NewTransactionHistoryRepo(r.SqlxDb))
}

func NewRepoManager(sqlxDb *sqlx.DB, gormDb *gorm.DB) RepoManager {
	return &repoManager{
		sqlxDb,
		gormDb,
	}
}
