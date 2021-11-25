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

func (u *User) SetFullName(name, lastname string) {
	u.Name = name
	u.Lastname = lastname
}

func (u *User) SetAge(age int) {
	u.Age = age
}

func (u *User) SetEmail(email string) {
	u.Email = email
}

func (u *User) SetPassword(password string) {
	u.Password = password
}
