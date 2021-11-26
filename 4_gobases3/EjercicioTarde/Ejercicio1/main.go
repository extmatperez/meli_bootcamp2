package main

import "fmt"

type usuario struct {
	Nombre     string
	Apellido   string
	Edad       int
	Correo     string
	Contrasena string
}

func setNombre(nombre string, u *usuario) {
	u.Nombre = nombre
}

func setApellido(apellido string, u *usuario) {
	u.Apellido = apellido
}

func setEdad(edad int, u *usuario) {
	u.Edad = edad
}

func setCorreo(correo string, u *usuario) {
	u.Correo = correo
}

func setContrasena(contrasena string, u *usuario) {
	u.Contrasena = contrasena
}

func main() {

	usuario1 := usuario{"Usuario1", "Apellido1", 18, "us1@gmail.com", "passus1"}
	fmt.Printf("\n%+v\n", usuario1)
	setNombre("CambioUsuario1", &usuario1)
	setApellido("CambioApellido", &usuario1)
	setEdad(19, &usuario1)
	setCorreo("cambiocorreous1@gmail.com", &usuario1)
	setContrasena("cambiopassus1", &usuario1)
	fmt.Printf("\n%+v\n", usuario1)

}
