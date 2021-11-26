package main

type usuario struct {
	Nombre     string
	Apellido   string
	Edad       int
	Correo     string
	Contrasena string
}

func cambiarNombre(u *usuario, nombre string) {
	u.Nombre = nombre
}

func cambiarEdad(u *usuario, edad int) {
	u.Edad = edad
}

func cambiarCorreo(u *usuario, correo string) {
	u.Correo = correo
}

func cambiarContrasena(u *usuario, contrasena string) {
	u.Contrasena = contrasena
}

func main() {

	us1 := usuario{Nombre: "pedro", Apellido: "perez", Edad: 20, Correo: "pedro@hotmail.com", Contrasena: "pedro123"}

	cambiarNombre(&us1, "jose")
	cambiarEdad(&us1, 21)
	cambiarCorreo(&us1, "jose@hotmail.com")
	cambiarContrasena(&us1, "jose123")
}
