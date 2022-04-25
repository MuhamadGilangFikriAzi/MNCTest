package model

type AdminRoom struct {
	Id          string
	Username    string
	Password    string
	Name        string
	PhoneNumber string `db:"phone_number"`
	IsActive    bool   `db:"is_active"`
}
