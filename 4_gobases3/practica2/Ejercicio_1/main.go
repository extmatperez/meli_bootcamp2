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
	updateName("Juan", "Soza", &user)
	updateAge(18, &user)
	updateEmail("Benjamin@gmail.com", &user)
	updatePassword("1234", &user)
	fmt.Println(user)
}
func updateName (name string, lastName string, u *Users) {
	u.firstName = name
	u.lastName = lastName
}
func updateAge (age int, u *Users){
	u.age = age
}
func updateEmail (email string, u *Users){
	u.email = email
}
func updatePassword (password string, u *Users){
	u.password = password
}
