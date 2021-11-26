package main

import "fmt"

type Usuarios struct {
	Nombre     string
	Edad       int
	Correo     string
	Contraseña string
}

func cambiarNombre(usuario *Usuarios, nombreNuevo string) {
	usuario.Nombre = nombreNuevo
}

func cambiarEdad(usuario *Usuarios, edadNueva int) {
	usuario.Edad = edadNueva
}

func cambiarCorreo(usuario *Usuarios, correoNuevo string) {
	usuario.Correo = correoNuevo
}

func cambiarContraseña(usuario *Usuarios, contraseñaNueva string) {
	usuario.Contraseña = contraseñaNueva
}

func main() {
	usuario1 := Usuarios{"Damian", 32, "damian.zamora@hotmail.com", "123456"}
	fmt.Println(usuario1)
	cambiarNombre(&usuario1, "Daniel")
	cambiarEdad(&usuario1, 35)
	cambiarCorreo(&usuario1, "dami.zamora@gmail.com")
	cambiarContraseña(&usuario1, "654321")
	fmt.Println(usuario1)
}
