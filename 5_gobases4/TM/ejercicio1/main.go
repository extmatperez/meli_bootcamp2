/*Una empresa de redes sociales requiere implementar una estructura usuario con funciones que vayan agregando información a la
estructura. Para optimizar y ahorrar memoria requieren que la estructura usuarios ocupe el mismo lugar en memoria para el main
del programa y para las funciones:
La estructura debe tener los campos: Nombre, Apellido, edad, correo y contraseña
Y deben implementarse las funciones:
cambiar nombre: me permite cambiar el nombre y apellido.
cambiar edad: me permite cambiar la edad.
cambiar correo: me permite cambiar el correo.
cambiar contraseña: me permite cambiar la contraseña.
*/
package main

import "fmt"

var lista []Usuario

type Usuario struct {
	Nombre, Apellido   string
	Edad               int
	Correo, Contraseña string
}

func cambiarNombre(user Usuario) {
	user.Nombre = ("Joselito")
	user.Apellido = ("Joselin")
	//lista = append(lista, user)

}

func cambiarEdad(user Usuario) {
	user.Edad = (50)
	//lista = append(lista, user)

}

func cambiarCorreo(user Usuario) {
	user.Correo = "eljoselito@gmail.com"
	//lista = append(lista, user)

}
func cambiarContraseña(user Usuario) {
	user.Contraseña = "contraseña"
	lista = append(lista, user)

}

func main() {
	user1 := Usuario{"Juan", "Perez", 23, "juanperez@gmail.com", "asd123"}
	user2 := Usuario{"Violeta", "Perez", 23, "violet@gmail.com", "asd123"}
	lista = append(lista, user1, user2)
	//lista[1] = Usuario{"Jose", "Perez", 23, "juanperez@gmail.com", "asd123"}
	fmt.Println(lista)

	cambiarNombre(user1)
	cambiarEdad(user1)
	cambiarCorreo(user1)
	cambiarContraseña(user1)
	fmt.Printf("\n%+v", lista)

}

//casi que esta bien ja ja
