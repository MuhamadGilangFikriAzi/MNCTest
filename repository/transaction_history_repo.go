package repository

import "github.com/jmoiron/sqlx"

type TransactionHistoryRepo interface {
	SaveHistoryLog(customerId string, activity string)
}

type transactionHistoryRepo struct {
	db *sqlx.DB
}

func (repo *transactionHistoryRepo) SaveHistoryLog(customerId string, activity string) {
	_, err := repo.db.Query("insert into history_customer(customer_id, activity) values($1, $2)", customerId, activity)
	if err != nil {
		panic(err.Error())
	}
}

func NewTransactionHistoryRepo(db *sqlx.DB) TransactionHistoryRepo {
	return &transactionHistoryRepo{db}
}
