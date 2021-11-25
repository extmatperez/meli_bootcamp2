package main

import "fmt"

type User struct {
	name     string
	age      int
	email    string
	password string
}

func (u *User) setName(name string) {
	u.name = name
}

func (u *User) setAge(age int) {
	u.age = age
}

func (u *User) setMail(email string) {
	u.email = email
}

func (u *User) setPassword(password string) {
	u.password = password
}

func main() {
	var u = User{}
	fmt.Println(u)
	u.setName("Federico Archuby")
	u.setAge(31)
	u.setMail("federico.archuby@mercadolibre.com")
	u.setPassword("federico")

	fmt.Println(u)
}
