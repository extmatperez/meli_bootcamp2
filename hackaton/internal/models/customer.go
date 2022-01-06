package models

type Customer struct {
	Id        int    `json:"id"`
	LastName  string `json:"lastName"`
	FirstName string `json:"firstName"`
	Condition string `json:"condition"`
}
