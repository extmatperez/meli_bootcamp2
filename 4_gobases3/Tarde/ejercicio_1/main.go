package main

import "fmt"

type usuario struct {
	Nombre     string
	Apellido   string
	edad       int
	correo     string
	contraseña string
}

func cambNombreApellido(u *usuario, nombre string, apellido string) {
	u.Nombre = nombre
	u.Apellido = apellido
}
func cambEdad(u *usuario, edad int) {
	u.edad = edad
}
func cambCorreo(u *usuario, correo string) {
	u.correo = correo
}
func cambContra(u *usuario, contra string) {
	(*u).contraseña = contra
}

func main() {
	u1 := usuario{
		Nombre:     "Diego",
		Apellido:   "Parra",
		edad:       25,
		correo:     "diego@correo.com",
		contraseña: "123",
	}

	fmt.Println("\nNombre antes:", u1.Nombre)
	fmt.Println("Apellido antes:", u1.Apellido)
	fmt.Println("edad antes:", u1.edad)
	fmt.Println("Correo antes:", u1.correo)
	fmt.Println("Contraseña antes:", u1.contraseña)

	cambNombreApellido(&u1, "Alejandro", "Parra")
	cambEdad(&u1, 15)
	cambCorreo(&u1, "nuevo@mail.com")
	cambContra(&u1, "000")

	fmt.Println("\n\nNombre despues:", u1.Nombre)
	fmt.Println("Apellido despues:", u1.Apellido)
	fmt.Println("edad despues:", u1.edad)
	fmt.Println("Correo despues:", u1.correo)
	fmt.Println("Contraseña despues:", u1.contraseña)

}
