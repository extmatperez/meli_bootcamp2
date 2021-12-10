package main

import (
	"fmt"
)

type Usuario struct {
	Nombre     string
	Apellido   string
	edad       int
	correo     string
	contraseña string
}

func CambiarName(user *Usuario) {
	var nuevoNombre, nuevoApellido string
	fmt.Println("ingrese el nuevo nombre a cambiar")
	fmt.Scanln(nuevoNombre)
	user.Nombre = "pipe"

	fmt.Println("ingrese el nuevo apellido")
	fmt.Scanln(nuevoApellido)
	user.Apellido = "montero"
}

func main() {
	User := Usuario{"Andres", "Pachon", 22, "andres@gmail.com", "123456"}
	CambiarName(&User)
	fmt.Println(User)
}
