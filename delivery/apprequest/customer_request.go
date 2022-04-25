package apprequest

type CustomerRequest struct {
	CustomerId     string `json:"customerId,omitempty"`
	Name           string `json:"name"`
	TypeCustomerId string `json:"type_customer_id"`
	Username       string `json:"username,omitempty"`
	Password       string `json:"password,omitempty"`
}
