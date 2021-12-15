package main

import "fmt"

type User struct {
	FName    string
	LName    string
	Age      int
	Email    string
	Password string
}

func (u *User) setUser(fname, lname, email, password string, age int) {
	u.FName = fname
	u.LName = lname
	u.Age = age
	u.Email = email
	u.Password = password
}

func (u *User) switchFName(fname string) {
	(*u).FName = fname
}

func (u *User) switchLName(lname string) {
	(*u).LName = lname
}

func (u *User) switchAge(age int) {
	(*u).Age = age
}

func (u *User) switchEmail(email string) {
	(*u).Email = email
}

func (u *User) switchPassword(password string) {
	(*u).Password = password

}

func main() {

	userStructure := User{}

	userStructure.switchFName("Facundo")
	userStructure.switchLName("Centeno")
	userStructure.switchAge(27)
	userStructure.switchEmail("facundocenteno@outlook.es")
	userStructure.switchPassword("*************")
	fmt.Println(userStructure)

}
