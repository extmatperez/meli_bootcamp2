package main

import "fmt"

type User struct {
	Name     string
	LName    string
	Age      int
	Email    string
	Password string
}

func setName(name, lname string, u *User) {
	(*u).Name = name
	(*u).LName = lname
}

func setAge(age int, u *User) {
	(*u).Age = age
}

func setEmail(email string, u *User) {
	(*u).Email = email
}

func setPassword(password string, u *User) {
	(*u).Password = password
}

func main() {
	newUser := User{Name: "Gabriel", LName: "Medina", Age: 47, Email: "gabmedina@gmail.com", Password: "flores"}

	fmt.Println("antes", newUser)

	setName("Pablo", "Lamponne", &newUser)
	setAge(46, &newUser)
	setEmail("pablinclavin@gmail.com", &newUser)
	setPassword("clavos", &newUser)

	fmt.Println("despues", newUser)
	fmt.Printf("Direccion de memoria: %p\n", &newUser)
	fmt.Printf("Valor que esta siendo apuntado: %v\n", newUser)
}
