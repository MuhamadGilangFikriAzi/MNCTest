package appresponse

type UserTransactionResonse struct {
	Name       string
	RoomNumber string `db:"room_number" json:"room_number"`
	Price      string
	OrderDate  string `db:"order_date" json:"order_date"`
	Status     string
}
