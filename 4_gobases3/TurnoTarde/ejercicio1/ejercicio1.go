package main

import (
	"fmt"
)

type Usuario struct {
	Nombre     string
	Edad       uint8
	Correo     string
	Contrasena string
}

func AgregarUsuario(usuarios *[]Usuario, usuario Usuario) {
	*usuarios = append(*usuarios, usuario)
	fmt.Println(*usuarios)
}

func CambiarNombre(usuario *Usuario, nombre string) {
	(*usuario).Nombre = nombre
}
func CambiarEdad(usuario *Usuario, edad uint8) {
	(*usuario).Edad = edad
}
func CambiarCorreo(usuario *Usuario, correo string) {
	(*usuario).Correo = correo
}
func CambiarContrasena(usuario *Usuario, contrasena string) {
	(*usuario).Contrasena = contrasena
}

func main() {

	usuario1 := Usuario{"Carlos Gonzalez", 58, "cgonzalez@gmail.com", "123456"}
	usuario2 := Usuario{"Maria Fernandez", 18, "mfernandez@gmail.com", "123456"}
	usuario3 := Usuario{"Sofia Martinez", 62, "smartinez85@gmail.com", "123456"}
	usuario4 := Usuario{"Juan Perez", 29, "jperez@gmail.com", "123456"}
	usuario5 := Usuario{"Jose Ruiz", 37, "jruiz@gmail.com", "123456"}

	var pUsuario *Usuario
	pUsuario = &usuario1
	CambiarNombre(pUsuario, "Felipe")
	CambiarEdad(pUsuario, 99)
	CambiarCorreo(pUsuario, "cgonzalez@hotmail.com")
	CambiarContrasena(pUsuario, "ahjsdhjkasd*&(&*(&*")

	usuarios := []Usuario{usuario1, usuario2, usuario3, usuario4, usuario5}
	fmt.Printf("%v\n", usuarios)
	usuario6 := Usuario{"Mario Alcazar", 27, "ma.785@gmail.com", "123456"}

	var pUsuarios *[]Usuario
	pUsuarios = &usuarios
	AgregarUsuario(pUsuarios, usuario6)

	fmt.Printf("%v\n", usuarios)
}
