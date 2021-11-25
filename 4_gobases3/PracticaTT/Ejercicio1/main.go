package main

import "fmt"

/*Una empresa de redes sociales requiere implementar una estructura usuario con funciones que vayan agregando
información a la estructura. Para optimizar y ahorrar memoria requieren que la estructura usuarios ocupe el mismo
lugar en memoria para el main del programa y para las funciones:
La estructura debe tener los campos: Nombre, Apellido, edad, correo y contraseña
Y deben implementarse las funciones:
cambiar nombre: me permite cambiar el nombre y apellido.
cambiar edad: me permite cambiar la edad.
cambiar correo: me permite cambiar el correo.
cambiar contraseña: me permite cambiar la contraseña.
*/

type usuario struct {
	nombre     string
	apellido   string
	edad       int
	correo     string
	contraseña string
}

func (u *usuario) cambiarNombre(nombre string, apellido string) {
	u.nombre = nombre
	u.apellido = apellido
}
func (u *usuario) cambiarEdad(edad int) {
	u.edad = edad
}
func (u *usuario) cambiarCorreo(correo string) {
	u.correo = correo
}
func (u *usuario) cambiarContraseña(contraseña string) {
	u.contraseña = contraseña
}

func main() {
	usuario1 := usuario{}

	fmt.Println("Usuario inicializado", usuario1)

	usuario1.cambiarNombre("Facundo", "Bouza")
	fmt.Println("Usuario con nombre", usuario1)

	usuario1.cambiarEdad(23)
	fmt.Println("Usuario con edad", usuario1)

	usuario1.cambiarCorreo("facubouza@gmail.com")
	fmt.Println("Usuario con mail", usuario1)

	usuario1.cambiarContraseña("asd123")
	fmt.Println("Usuario con contraseña", usuario1)

}
