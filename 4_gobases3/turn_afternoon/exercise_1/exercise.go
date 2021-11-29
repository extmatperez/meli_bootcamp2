package main

import "fmt"

func main() {

	user := User{}

	var us = new(User)

	//us.SetNameAndLastName("Maru", "Davalos")

	//fmt.Println(us)

	us = &user

	user.SetNameAndLastName("Dig", "Davila")
	user.SetAge(26)
	user.SetEmail("dig@example.com")
	user.SetPassword("123456.")

	fmt.Printf("The main user is: %v\n", user)
	fmt.Printf("The main user is: %v\n", *us)

	SetNameAndLastName2("Matilda", "Rolan", &user)
	SetAge2(35, &user)
	SetEmail2("matilda@example.com", &user)
	SetPassword2("password.", &user)

	fmt.Printf("The main user is: %v\n", user)

}

type User struct {
	Name     string
	LastName string
	Age      int
	Email    string
	Password string
}

func (u *User) SetNameAndLastName(name string, lastName string) {
	u.Name = name
	u.LastName = lastName
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

func SetNameAndLastName2(name string, lastName string, u *User) {
	(*u).Name = name
	(*u).LastName = lastName
}

func SetAge2(age int, u *User) {
	(*u).Age = age
}

func SetEmail2(email string, u *User) {
	(*u).Email = email
}

func SetPassword2(password string, u *User) {
	(*u).Password = password
}
