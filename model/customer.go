package model

import (
	"github.com/jackc/pgtype"
)

type Customer struct {
	Id          string
	Name        string
	Address     string
	Ktp         string      `db:"ktpid"`
	PhoneNumber string      `db:"phone_number"`
	StartRentAt pgtype.Date `db:"start_rent_at"`
	EndRentAt   pgtype.Date `db:"end_rent_at"`
}
