package main

import (
	"fmt"
)

// Red social
type Usuario struct {
	Nombre string
	Apellido string
	Edad int
	Correo string
	Contraseña string
}

func updateNombre(nombre, apellido string, u *Usuario) {
	u.Nombre = nombre
	u.Apellido = apellido
}

func (u *Usuario) updateEdad(edad int) {
	u.Edad = edad
}

func (u *Usuario) updateCorreo(correo string) {
	u.Correo = correo
}

func updateContrasenia(contraseña string, u *Usuario) {
	(*u).Contraseña = contraseña
}

func main() {

	user1 := Usuario{"Juan", "Snow", 25, "as@gmail.com", "123"}

	//fmt.Printf("%s\n", user1.Nombre)

	updateNombre("Jhon", "Snow", &user1)

	//fmt.Printf("%s\n", user1.Nombre)

	updateContrasenia("ffff", &user1)

	fmt.Printf("%+v\n", user1)
}