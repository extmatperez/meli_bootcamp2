package main

import "fmt"

// Ejercicio 1 - Red social
// Una empresa de redes sociales requiere implementar una estructura usuario con funciones que vayan agregando información a la estructura. Para optimizar y ahorrar memoria requieren que la estructura usuarios ocupe el mismo lugar en memoria para el main del programa y para las funciones:
// La estructura debe tener los campos: Nombre, Apellido, edad, correo y contraseña
// Y deben implementarse las funciones:
// cambiar nombre: me permite cambiar el nombre y apellido.
// cambiar edad: me permite cambiar la edad.
// cambiar correo: me permite cambiar el correo.
// cambiar contraseña: me permite cambiar la contraseña.


	type User struct {
		Name 		string
		Surname 	string
		Age 		int
		Mail 		string
		Password 	string
	}

	func (u *User) change_name(new_name, new_surname string) {
		u.Name = new_name
		u.Surname = new_surname
	}

	func (u *User) change_age(new_age int) {
		u.Age = new_age
	}

	func (u *User) change_mail(new_mail string) {
		u.Mail = new_mail
	}

	func (u *User) change_password(new_password string) {
		u.Password = new_password
	}


func main () {


	new_user1 := User{Name: "Pepe", Surname: "Perez", Age: 34, Mail: "pepeperez@mail.com", Password: "1234"}

	fmt.Println(new_user1)

	new_user1.change_name("Dario", "Gonzales")

	fmt.Println(new_user1)




}