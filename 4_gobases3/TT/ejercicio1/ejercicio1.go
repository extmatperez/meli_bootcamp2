package main

/*
Una empresa de redes sociales requiere implementar una estructura usuario con funciones que vayan agregando información a la estructura. Para optimizar y ahorrar memoria requieren que la estructura usuarios ocupe el mismo lugar en memoria para el main del programa y para las funciones:
La estructura debe tener los campos: Nombre, Apellido, edad, correo y contraseña Y deben implementarse las funciones:
- cambiar nombre: me permite cambiar el nombre y apellido.
- cambiar edad: me permite cambiar la edad.
- cambiar correo: me permite cambiar el correo.
- cambiar contraseña: me permite cambiar la contraseña.
*/

import (
	"fmt"
)

type Usuario struct {
	Nombre     string
	Apellido   string
	Edad       int
	Correo     string
	Contraseña string
}

func (user *Usuario) cambiarNombre(name string, last string) {
	user.Nombre = name
	user.Apellido = last
}
func (user *Usuario) cambiarEdad(age int) {
	user.Edad = age
}
func (user *Usuario) cambiarCorreo(mail string) {
	user.Correo = mail
}
func cambiarPass(user *Usuario, pass string) {
	(*user).Contraseña = pass
}

func main() {
	fmt.Println("Ejercicio1")
	var user Usuario
	user.Nombre = "Juan"
	user.Apellido = "Perez"
	user.Edad = 23
	user.Correo = "as@hotmail.com"
	user.Contraseña = "12345"

	fmt.Println("Usuario Inicial: ", user)
	user.cambiarNombre("Pepe", "Argento")
	user.cambiarEdad(30)
	user.cambiarCorreo("dddd@hotmail.com")
	cambiarPass(&user, "mamamia")
	fmt.Println("Usuario final: ", user)
}
