package main

import "fmt"

type Usuario struct {
	Nombre      string
	Apellido    string
	Edad        int
	Correo      string
	Contrasenia string
}

func cambiarNomyApe(newNombre, newApellido string, u *Usuario) {
	(*u).Nombre = newNombre
	(*u).Apellido = newApellido
}

func cambiarEdad(newEdad int, u *Usuario) {
	(*u).Edad = newEdad
}

func cambiarCorreo(newCorreo string, u *Usuario) {
	(*u).Correo = newCorreo
}

func cambiarPass(newPass string, u *Usuario) {
	(*u).Contrasenia = newPass
}

func main() {
	var p1 *Usuario
	//var pn *int
	//num := 5
	u1 := Usuario{"Nicolas", "Aponte", 23, "nico@correo.com", "sdfsd"}
	p1 = &u1
	//pn = &num
	//fmt.Println(*pn, " ", &num, " ", pn, " ", num, " ", &pn)
	//fmt.Println(&u1)
	//fmt.Println((*p1), " ", &u1, " ", p1, " ", u1, " ", &p1)
	cambiarNomyApe("Pedro", "Pepe", p1)
	cambiarEdad(32, p1)
	cambiarCorreo("@gmail.com", p1)
	cambiarPass("----", p1)
	fmt.Println(u1)

}
