package apprequest

type TransactionRequest struct {
	CustomerId string `json:"customer_id"`
	RoomNumber string `json:"room_number"`
}

type TransactionUpdateRequest struct {
	CustomerId string `json:"customer_id" binding:"required"`
	Status     int    `json:"status"`
}
