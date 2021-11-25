package main

import "fmt"

type User struct {
	name string
	lastname string
	age int
	email string
	password string
}

func (u *User) setName(name, lastname string) {
	u.name = name
	u.lastname = lastname
}

func (u *User) setAge(age int) {
	u.age = age
}

func (u *User) setEmail(email string) {
	u.email = email
}

func (u *User) setPassword(password string) {
	u.password = password
}

func main() {
	user := User{}
	user.setName("Juan", "Perez")
	user.setAge(30)
	user.setEmail("juan@mail.com")
	user.setPassword("12345")
	fmt.Println(user)
}