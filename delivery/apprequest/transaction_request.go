package apprequest

type TransactionRequest struct {
	SenderId  string  `json:"sender_id"`
	ReceiveId string  `json:"receive_id"`
	Amount    float64 `json:"amount"`
}
