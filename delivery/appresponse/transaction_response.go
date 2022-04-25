package appresponse

import "time"

type TransactionResponse struct {
	SenderName      string    `json:"sender_name,omitempty"`
	ReceiverName    string    `json:"receiver_name,omitempty"`
	Amount          string    `json:"amount,omitempty"`
	TransactionDate time.Time `json:"transaction_date,omitempty"`
}
