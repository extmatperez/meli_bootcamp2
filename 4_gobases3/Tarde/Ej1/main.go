/*
?Ejercicio 1 - Red social
Una empresa de redes sociales requiere implementar una estructura usuario
con funciones que vayan agregando información a la estructura. Para optimizar
y ahorrar memoria requieren que la estructura usuarios ocupe el mismo lugar
en memoria para el main del programa y para las funciones:
La estructura debe tener los campos: Nombre, Apellido, edad, correo y
contraseña
Y deben implementarse las funciones:
cambiar nombre: me permite cambiar el nombre y apellido.
cambiar edad: me permite cambiar la edad.
cambiar correo: me permite cambiar el correo.
cambiar contraseña: me permite cambiar la contraseña.

*/

package main

import "fmt"

type Usuario struct {
	Nombre     string
	Apellido   string
	Edad       int
	Correo     string
	Contraseña string
}

func setNombre(nuevoNombre, nuevoApellido string, u *Usuario) {
	(*u).Nombre = nuevoNombre
	(*u).Apellido = nuevoApellido
}

func setEdad(nuevaEdad int, u *Usuario) {
	(*u).Edad = nuevaEdad
}

func setCorreo(nuevoCorreo string, u *Usuario) {
	(*u).Correo = nuevoCorreo
}

func setContraseña(nuevaContraseña string, u *Usuario) {
	(*u).Contraseña = nuevaContraseña
}

func main() {

	usuario1 := Usuario{"Nico", "Arguello", 36, "nico@gmail.com", "4567"}
	fmt.Println(usuario1)

	setNombre("Cele", "Gonzalez", &usuario1)
	setEdad(23, &usuario1)
	setCorreo("cele@gmail.com", &usuario1)
	setContraseña("1234", &usuario1)

	fmt.Println(usuario1)
}
