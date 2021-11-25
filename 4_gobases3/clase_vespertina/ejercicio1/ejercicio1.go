package main

import (
	"fmt"
)

func main() {

	usuario1 := Usuario{"Carlos Gianni", 58, "cgianni@gmail.com", "123456"}
	usuario2 := Usuario{"Veronica Estevez", 18, "vestevez@gmail.com", "123456"}
	usuario3 := Usuario{"Maria Fontoura", 62, "mfra85@gmail.com", "123456"}
	usuario4 := Usuario{"Mateo Perez", 29, "mperez5@gmail.com", "123456"}
	usuario5 := Usuario{"Paola Pocaz", 37, "paola.pocaz19@gmail.com", "123456"}

	var pUsuario *Usuario
	pUsuario = &usuario1
	CambiarNombre(pUsuario, "Felipe")
	CambiarEdad(pUsuario, 99)
	CambiarCorreo(pUsuario, "fgianni@hotmail.com")
	CambiarContrasena(pUsuario, "ahjsdhjkasd*&(&*(&*(&jkasdhkajsdhjkJKHJKADSK")

	usuarios := []Usuario{usuario1, usuario2, usuario3, usuario4, usuario5}
	fmt.Printf("%v\n", usuarios)
	usuario6 := Usuario{"Mario Alcazar", 27, "ma.785@gmail.com", "123456"}

	var pUsuarios *[]Usuario
	pUsuarios = &usuarios
	AgregarUsuario(pUsuarios, usuario6)

	fmt.Printf("%v\n", usuarios)
}

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
