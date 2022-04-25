package appresponse

type CustomerResponse struct {
	CustomerId      string               `json:"customer_id,omitempty"`
	Name            string               `json:"name"`
	Balance         float64              `json:"balance"`
	BalanceFormated string               `json:"balance_formated"`
	Username        string               `json:"username,omitempty"`
	TypeCustomer    TypeCustomerResponse `json:"type_customer"`
}
