package main

import (
	"fmt"
	
)

type Usuario struct {
	Nombre,Apellido,Correo,Contraseña string 
	Edad int 

}

func (user *Usuario) SetName(nombre, apellido string){
	user.Apellido=apellido
	user.Nombre=nombre
}
func (user *Usuario) SetEdad(edad int){
	user.Edad= edad
}
func (user *Usuario) SetCorreo(correo string){
	user.Correo= correo
}
func (user *Usuario) SetPasswrod(contrasena string){
	user.Contraseña= contrasena
}





func main() {
	usuario := Usuario{Nombre: "Pepe",Apellido: "Castro",Edad: 15,Correo: "safñjasf@gmail.com", Contraseña: "123456"}
	fmt.Println("Usuario original",usuario)

	usuario.SetName("pancho","palacio")
	usuario.SetCorreo("luis@yahoo.com.ar")
	usuario.SetEdad(88)
	usuario.SetPasswrod("654321")
	fmt.Println("Usuario cambiado",usuario)

 }


