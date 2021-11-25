package main

import "fmt"

func main() {
	user := User{
		Name:     "",
		Lastname: "",
		Age:      0,
		Email:    "",
		Password: "",
	}

	fmt.Println(user)
	fmt.Println()

	user.SetFullName("Matias", "Ziliotto")
	user.SetAge(24)
	user.SetEmail("matias.ziliotto@mercadolibre.com")
	user.SetPassword("myPassword")

	fmt.Println(user)
	fmt.Println()
}

type User struct {
	Name     string
	Lastname string
	Age      int
	Email    string
	Password string
}

func (p *User) SetFullName(name, lastname string) {
	p.Name = name
	p.Lastname = lastname
}

func (p *User) SetAge(age int) {
	p.Age = age
}

func (p *User) SetEmail(email string) {
	p.Email = email
}

func (p *User) SetPassword(password string) {
	p.Password = password
}
