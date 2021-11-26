package main

import "fmt"

// cambiar nombre: me permite cambiar el nombre y apellido.
// cambiar edad: me permite cambiar la edad.
// cambiar correo: me permite cambiar el correo.
// cambiar contraseña: me permite cambiar la contraseña.

type UsuarioRed struct {
	Nombre, Apellido, Correo, Contrasenia string
	Edad                                  int
}

func (u *UsuarioRed) cambiarNombre(nuevoNombre string, nuevoApellido string) {
	u.Nombre = nuevoNombre
	u.Apellido = nuevoApellido
}
func (u *UsuarioRed) cambiarEdad(nuevaEdad int) {
	u.Edad = nuevaEdad
}
func (u *UsuarioRed) cambiarCorreo(nuevoCorreo string) {
	u.Correo = nuevoCorreo
}
func (u *UsuarioRed) cambiarContrasenia(nuevaContrasenia string) {
	u.Contrasenia = nuevaContrasenia
}
func main() {
	u := UsuarioRed{"Juan Pablo", "Pescie", "jpescie@meli.com", "Password", 20}
	fmt.Println(u)
	u.cambiarNombre("Nuevo nombre", "Apellido")
	u.cambiarEdad(56)
	u.cambiarCorreo("correo@mlibre.com")
	u.cambiarContrasenia("Contras")
	fmt.Println(u)
}
