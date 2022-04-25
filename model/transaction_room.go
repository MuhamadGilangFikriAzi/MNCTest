package model

import (
	"database/sql"
	"github.com/jackc/pgtype"
)

type TransactionRoom struct {
	Id         string
	AdminId    string         `db:"admin_id"`
	CusomerId  string         `db:"customer_id"`
	RoomNumber string         `db:"room_number"`
	CreatedAt  pgtype.Date    `db:"created_at"`
	UpdatedAt  sql.NullString `db:"updated_at"`
	Status     int            // 1 = pending, 2 = not paid, 3 = finish
}
