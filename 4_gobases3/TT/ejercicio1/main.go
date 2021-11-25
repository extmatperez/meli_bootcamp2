package main

import "fmt"

type Usuario struct {
	Nombre     string
	Apellido   string
	Edad       int
	Correo     string
	Contrasena string
}

func main() {
	user := Usuario{"Seba", "Chiappa", 26, "correo@mail.com", "p4ssw0rd"}
	fmt.Println(user)
	CambiarNombre(&user, "Sebastian", "Ch")
	CambiarEdad(&user, 30)
	CambiarCorreo(&user, "mail@gmail.com")
	CambiarContrasenia(&user, "1234")
	fmt.Println(user)

}
func CambiarNombre(usr *Usuario, nombre string, apellido string) {
	usr.Nombre = nombre
	usr.Apellido = apellido
}

func CambiarEdad(usr *Usuario, edad int) {
	usr.Edad = edad
}

func CambiarCorreo(usr *Usuario, correo string) {
	usr.Correo = correo
}

func CambiarContrasenia(usr *Usuario, contasenia string) {
	usr.Contrasena = contasenia
}
