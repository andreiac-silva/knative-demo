package models

type Customer struct {
	Cpf string `json:"cpf"`
}
type Items struct {
	ID          int     `json:"id"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
}
type Purchase struct {
	Customer Customer `json:"customer"`
	Items    []Items  `json:"items"`
}
