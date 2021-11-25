/*
Una empresa de redes sociales requiere implementar una estructura usuario con funciones
que vayan agregando información a la estructura. Para optimizar y ahorrar memoria requieren
que la estructura usuarios ocupe el mismo lugar en memoria para el main del programa y para las funciones:

La estructura debe tener los campos: Nombre, Apellido, edad, correo y contraseña

Y deben implementarse las funciones:
cambiar nombre: me permite cambiar el nombre y apellido.
cambiar edad: me permite cambiar la edad.
cambiar correo: me permite cambiar el correo.
cambiar contraseña: me permite cambiar la contraseña.
*/

package main

import "fmt"

type usuario struct {
	nombre   string
	apellido string
	edad     int
	correo   string
	password string
}

func cambiarNombre(nuevoNombre string, u *usuario) {
	u.nombre = nuevoNombre
}

func cambiarEdad(nuevaEdad int, u *usuario) {
	u.edad = nuevaEdad
}

func cambiarCorreo(nuevoCorreo string, u *usuario) {
	u.correo = nuevoCorreo
}

func cambiarPassword(nuevaPassword string, u *usuario) {
	u.password = nuevaPassword
}

func main() {

	u := usuario{"Benjamin", "Conti", 25, "contibenjamin@gmail.com", "hola"}

	fmt.Println(u)

	cambiarNombre("Pepe", &u)
	cambiarEdad(99, &u)
	cambiarCorreo("pepito@mercadolibre.com", &u)
	cambiarPassword("holax", &u)

	fmt.Println(u)
}
