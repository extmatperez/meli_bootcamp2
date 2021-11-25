package main

import "fmt"

type usuario struct {
	Nombre     string
	Apellido   string
	Edad       string
	Correo     string
	Contraseña string
}

func (u *usuario) actualizar(campo string, data string) {
	switch campo {
	case "Nombre":
		u.Nombre = data
	case "Apellido":
		u.Apellido = data
	case "Edad":
		u.Edad = data
	case "Correo":
		u.Correo = data
	case "Contraseña":
		u.Contraseña = data
	default:
		fmt.Println("Usuario no valido")
	}
}

func main() {

	fmt.Println("Bienvenidos al ejercicio 1")

	usuario1 := usuario{"Patricio", "Pallua", "23", "Mail@mail.com", "contraseña"}

	fmt.Println("El usuario1 es", usuario1)

	fmt.Println("Probamos actualizar el apellido")
	usuario1.actualizar("Apellido", "Pessagno")
	fmt.Println("El usuario cambio su apellido", usuario1)

	fmt.Println("Probamos actualizar la contraseña")
	usuario1.actualizar("Contraseña", "nuevaPassword")
	fmt.Println("Usuario1 cambio su contraseña", usuario1)

}
