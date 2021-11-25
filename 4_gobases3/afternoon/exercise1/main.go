package main

import "fmt"

type Usuarios struct {
	Nombre     string `json:"nombre"`
	Apellido   string `json:"apellido"`
	Edad       int    `json:"edad"`
	Correo     string `json:"correo"`
	Contrasena string `json:"contrasena"`
}

func (u *Usuarios) cambiarNombre() {
	var nombre string
	var apellido string
	fmt.Println("Ingrese el nuevo nombre: ")
	fmt.Scanln(&nombre)
	fmt.Println("Ingrese el nuevo apellido: ")
	fmt.Scanln(&apellido)
	u.Nombre = nombre
	u.Apellido = apellido
}
func (u *Usuarios) cambiarEdad() {
	var edad int
	fmt.Println("Ingrese la nueva edad: ")
	fmt.Scanln(&edad)
	u.Edad = edad
}
func (u *Usuarios) cambiarCorreo() {
	var correo string
	fmt.Println("Ingrese el nuevo correo: ")
	fmt.Scanln(&correo)
	u.Correo = correo
}
func (u *Usuarios) cambiarContrasena() {
	var contrasena string
	fmt.Println("Ingrese la nueva contraseña: ")
	fmt.Scanln(&contrasena)
	u.Contrasena = contrasena
}

func main() {
	/*
		Una empresa de redes sociales requiere implementar una estructura usuario con funciones que vayan agregando información a la estructura. Para optimizar y ahorrar memoria requieren que la estructura usuarios ocupe el mismo lugar en memoria para el main del programa y para las funciones:
		La estructura debe tener los campos: Nombre, Apellido, edad, correo y contraseña
		Y deben implementarse las funciones:
		cambiar nombre: me permite cambiar el nombre y apellido.
		cambiar edad: me permite cambiar la edad.
		cambiar correo: me permite cambiar el correo.
		cambiar contraseña: me permite cambiar la contraseña.
	*/
	usuario1 := Usuarios{"Ariel", "Romero", 34, "ariel@g.com", "123"}
	usuario2 := Usuarios{"Pepe", "Rom", 34, "pepe@g.com", "321"}

	fmt.Println("Usuario 1: ", usuario1)
	fmt.Println("Usuario 2: ", usuario2)

	usuario1.cambiarNombre()
	usuario1.cambiarEdad()
	usuario1.cambiarCorreo()
	usuario1.cambiarContrasena()

	usuario2.cambiarNombre()

	fmt.Println("Usuario 1: ", usuario1)
	fmt.Println("Usuario 2: ", usuario2)
}
