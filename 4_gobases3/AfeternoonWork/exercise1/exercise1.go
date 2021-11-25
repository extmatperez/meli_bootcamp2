/*
Una empresa de redes sociales requiere implementar una estructura usuario con funciones que vayan agregando
información a la estructura. Para optimizar y ahorrar memoria requieren que la estructura usuarios ocupe el mismo
lugar en memoria para el main del programa y para las funciones:

- La estructura debe tener los campos: Nombre, Apellido, edad, correo y contraseña

Y deben implementarse las funciones:

- cambiar nombre: me permite cambiar el nombre y apellido.
- cambiar edad: me permite cambiar la edad.
- cambiar correo: me permite cambiar el correo.
- cambiar contraseña: me permite cambiar la contraseña.
*/

package main

import (
	"fmt"
)

type Users struct {
	Name     string
	LastName string
	Age      int
	Email    string
	Password string
}

func (user *Users) changeName(newName string) {
	user.Name = newName
}
func (user *Users) changeAge(newAge int) {
	user.Age = newAge
}
func (user *Users) changeEmail(newEmail string) {
	user.Email = newEmail
}
func (user *Users) changePassword(newPassword string) {
	user.Password = newPassword
}

func main() {
	usuario1 := Users{"Jose", "Rios", 28, "1111111", "passJose"}
	fmt.Println(usuario1)
	usuario1.changeName("Juan")
	usuario1.changeAge(32)
	usuario1.changeEmail("joserios@mercadolibre.cl")
	usuario1.changePassword("123OLVIDARpass")
	fmt.Println(usuario1)

}
