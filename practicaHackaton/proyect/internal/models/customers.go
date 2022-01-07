package models

type Customer struct {
	ID         int    `json:"id"`
	Last_Name  string `json:"last_name"`
	First_Name string `json:"first_name"`
	Condition  string `json:"condition"`
}

type RequestCustomer struct {
	Last_Name  string `json:"last_name"`
	First_Name string `json:"first_name"`
	Condition  string `json:"condition"`
}
