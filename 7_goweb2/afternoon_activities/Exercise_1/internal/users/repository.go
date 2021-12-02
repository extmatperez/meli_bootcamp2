package main

type Users struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
	Height    int    `json:"height"`
	Active    bool   `json:"active"`
	Date      string `json:"date"`
}

/* type Repository interface {
	get_users() ([]Users, error)
	post_users() (Users, error)
}

type Repository struct {}

var user []Users
var last_id int

func new_repository() Repository{
	return &repository{}
} */
