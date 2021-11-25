package main

import "fmt"

type Usuario struct {
	Nombre     string
	Apellido   string
	Edad       int
	Correo     string
	Contraseña string
}

func main() {
	usuario1 := Usuario{"Matias", "De Bonis", 20, "aaa", "bbb"}
	var puntero *Usuario = &usuario1
	cambiarNombreApellido(puntero, "Saitam", "Sinob Ed")
	cambiarEdad(puntero, -20)
	cambiarCorreo(puntero, "correo@mail.com")
	cambiarContraseña(puntero, "Contraseña")
	fmt.Println(usuario1)
}

func cambiarNombreApellido(usuario *Usuario, nuevoNombre, nuevoApellido string) {
	usuario.Nombre = nuevoNombre
	usuario.Apellido = nuevoApellido
}

func cambiarEdad(usuario *Usuario, nuevaEdad int) {
	usuario.Edad = nuevaEdad
}

func cambiarCorreo(usuario *Usuario, nuevoCorreo string) {
	usuario.Correo = nuevoCorreo
}

func cambiarContraseña(usuario *Usuario, nuevaContraseña string) {
	usuario.Contraseña = nuevaContraseña
}
