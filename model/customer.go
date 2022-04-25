package model

type Customer struct {
	Id             string  `db:"id"`
	Name           string  `db:"name"`
	TypeCustomerID string  `db:"type_customer_id"`
	Balance        float64 `db:"balance"`
	Username       string  `db:"username"`
	Password       string  `db:"password"`
}
