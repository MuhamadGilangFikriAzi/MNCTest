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

type Transaction interface {
	BalanceTransfer(data apprequest.TransactionRequest) (appresponse.TransactionResponse, string, error)
}

type transaction struct {
	db      *sqlx.DB
	history TransactionHistoryRepo
}

func (t *transaction) BalanceTransfer(data apprequest.TransactionRequest) (appresponse.TransactionResponse, string, error) {
	var dataSender model.Customer
	var dataReciver model.Customer
	var dataReturn appresponse.TransactionResponse
	tx := t.db.MustBegin()
	fmt.Println(data)
	errSender := tx.Get(&dataSender, "select id, name, balance, type_customer_id from customer where id = $1 and is_register = true and deleted_at is null", data.SenderId)
	if errSender != nil {
		return dataReturn, "Sender not found", errSender
	}
	errReciver := tx.Get(&dataReciver, "select id, name, balance, type_customer_id from customer where id = $1 and is_register = true and deleted_at is null", data.ReceiveId)
	if errReciver != nil {
		return dataReturn, "Receiver not found", errReciver
	}
	if dataSender.Balance < data.Amount {
		return dataReturn, "The balance is not sufficient", nil
	}
	dataSender.Balance -= data.Amount
	dataReciver.Balance += data.Amount
	tx.MustExec("update customer set balance = $1 where id = $2", dataSender.Balance, dataSender.Id)
	tx.MustExec("update customer set balance = $1 where id = $2", dataReciver.Balance, dataReciver.Id)
	tx.MustExec("insert into transaction_transfer(sender_id, receiver_id, amount) values($1, $2, $3)", dataSender.Id, dataReciver.Id, data.Amount)
	err := tx.Commit()
	if err != nil {
		return dataReturn, "Failed", err
	}
	dataReturn.TransactionDate = time.Now()
	dataReturn.ReceiverName = dataReciver.Name
	dataReturn.SenderName = dataSender.Name
	dataReturn.Amount = utility.CunrrencyFormat("Rp.", data.Amount)

	t.history.SaveHistoryLog(dataSender.Id, fmt.Sprintf("%s transfer %s to %s ", dataSender.Name, dataReturn.Amount, dataReciver.Name))
	return dataReturn, "Succes Transfer", nil
}

func NewTransaction(db *sqlx.DB, repoHisotry TransactionHistoryRepo) Transaction {
	return &transaction{db, repoHisotry}
}
