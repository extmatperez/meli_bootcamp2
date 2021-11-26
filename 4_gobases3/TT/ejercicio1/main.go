package main

import "fmt"

type Usuario struct {
	Nombre     string
	Apellido   string
	Edad       int
	Correo     string
	Contraseña string
}

func (u *Usuario) cambiar_nombre(nom string, ape string) {
	u.Nombre = nom
	u.Apellido = ape
}

func (u *Usuario) cambiar_edad(edad int) {
	u.Edad = edad
}

func (u *Usuario) cambiar_correo(email string) {
	u.Correo = email
}

func (u *Usuario) cambiar_contraseña(contra string) {
	u.Contraseña = contra
}

func main() {
	usuario := Usuario{}
	usuario.cambiar_nombre("diego", "maradona")
	usuario.cambiar_edad(24)
	usuario.cambiar_correo("ivan@gmail.com")
	usuario.cambiar_contraseña("contra1234")
	fmt.Print(usuario)
	usuario.cambiar_correo("hola@live.com")
	fmt.Print("\n", usuario, "\n")
}
