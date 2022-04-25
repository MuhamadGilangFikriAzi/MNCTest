package model

type BoardingRoom struct {
	RoomNumber  string `db:"room_number"`
	IsAvailable bool   `db:"is_available"`
	Price       int
	Facility    string
}
