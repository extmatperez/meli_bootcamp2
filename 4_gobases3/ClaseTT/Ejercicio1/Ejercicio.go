package main

import "fmt"

type usuario struct {
	Nombre     string
	Apellido   string
	Edad       int
	Correo     string
	Contraseña string
}

func (u *usuario) cambiarNombre(nombre, apellido string) {
	(*u).Nombre = nombre
	(*u).Apellido = apellido
}
func (u *usuario) cambiarCorreo(correo string) {
	(*u).Correo = correo
}
func (u *usuario) cambiarEdad(edad int) {
	(*u).Edad = edad
}
func (u *usuario) cambiarContrasenia(pass string) {
	(*u).Contraseña = pass
}
func main() {
	estructUser := usuario{}

	fmt.Println(estructUser)
	estructUser.cambiarNombre("Nahuel", "Scerca")
	estructUser.cambiarContrasenia("1234")
	estructUser.cambiarEdad(24)
	estructUser.cambiarCorreo("nahuel@gmail.com")
	fmt.Println(estructUser)

	estructUser2 := usuario{}
	fmt.Println(estructUser2)
	estructUser2.cambiarNombre("asdasdasd", "asdasd")
	estructUser2.cambiarContrasenia("asdasd")
	estructUser2.cambiarEdad(24)
	estructUser2.cambiarCorreo("nahuel@gmail.com")
	fmt.Println(estructUser2)
	fmt.Println(estructUser)

}
