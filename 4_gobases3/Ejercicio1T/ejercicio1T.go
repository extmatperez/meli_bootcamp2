package main

import "fmt"

type Usuario struct {
	Nombre, Apellido   string
	Edad               int
	Correo, Contrasena string
}

func (u *Usuario) setNombre(nombre string) {
	u.Nombre = nombre
}
func (u *Usuario) setEdad(edad int) {
	u.Edad = edad
}
func (u *Usuario) setCorreo(correo string) {
	u.Correo = correo
}
func (u *Usuario) setContrasena(contra string) {
	u.Contrasena = contra
}

// para hacerlo con funciones
func setConFunc(pass string, u *Usuario) {
	(*u).Contrasena = pass
}
func main() {

	usuario1 := Usuario{"Walter", "Castillo", 26, "walter@walter.com", "12345"}
	var pUsuario1 *Usuario

	fmt.Println()
	fmt.Println(usuario1, &pUsuario1)

	usuario1.setNombre("Terwal")
	usuario1.setEdad(27)
	usuario1.setCorreo("Terwal@terwal.com")
	usuario1.setContrasena("4321")

	fmt.Println(usuario1, &pUsuario1)

	setConFunc("12345678", &usuario1)
	fmt.Println(usuario1, &pUsuario1)

	fmt.Println()
}
