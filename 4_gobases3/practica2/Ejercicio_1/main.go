package main

import "fmt"

type Users struct {
	firstName	string
	lastName	string
	age			int
	email		string
	password	string
}

func main() {
	/*
	Una empresa de redes sociales requiere implementar una estructura usuario con funciones que vayan
	agregando información a la estructura. Para optimizar y ahorrar memoria requieren que la estructura
	usuarios ocupe el mismo lugar en memoria para el main del programa y para las funciones:
	La estructura debe tener los campos: Nombre, Apellido, edad, correo y contraseña
	Y deben implementarse las funciones:
	cambiar nombre: me permite cambiar el nombre y apellido.
		cambiar edad: me permite cambiar la edad.
		cambiar correo: me permite cambiar el correo.
		cambiar contraseña: me permite cambiar la contraseña.
	*/
	//setting User
	user := Users{}
	user.setFirstName()
	user.setLastName()
	user.setAge()
	user.setEmail()
	user.setPassword()
	fmt.Println(user)
}

func (u *Users) setFirstName() {
	var name string
	fmt.Println("Porfavor Ingrese el nombre: ")
	fmt.Scan(&name)
	u.firstName = name
}
func (u *Users) setLastName() {
	var lastName string
	fmt.Println("Porfavor Ingrese el apellido: ")
	fmt.Scan(&lastName)
	u.lastName = lastName
}
func (u *Users) setAge() {
	var age int
	fmt.Println("Porfavor Ingrese la Edad: ")
	fmt.Scan(&age)
	u.age = age
}
func (u *Users) setEmail() {
	var email string
	fmt.Println("Porfavor Ingrese la direccion de correo: ")
	fmt.Scan(&email)
	u.email = email
}
func (u *Users) setPassword() {
	var password string
	fmt.Println("Porfavor Ingrese la Contraseña: ")
	fmt.Scan(&password)
	u.password = password
}