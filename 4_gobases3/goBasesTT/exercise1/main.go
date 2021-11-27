package main

import "fmt"

const (
	nombre     = "nombre"
	apellido   = "apellido"
	correo     = "correo"
	contraseña = "contraseña"
	edad       = "edad"
)

func main() {
	var nombre, apellido, correo, contraseña, edad string
	//var edad int
	var cambio, cambio2, valor2 string
	fmt.Println("Ejercicio 1")
	fmt.Println("introduce los siguientes datos: nombre, apellido, correo, contraseña y edad")
	fmt.Scanf("%s", &nombre)
	fmt.Scanf("%s", &apellido)
	fmt.Scanf("%s", &correo)
	fmt.Scanf("%s", &contraseña)
	fmt.Scanf("%v", &edad)

	p1 := Persona{nombre, apellido, correo, contraseña, edad}
	p1.Print()

	fmt.Print("\nqueres realizar algun cambio?")
	fmt.Scanf("%s", &cambio)
	if cambio == "si" {
		fmt.Print("que deseas cambiar?")
		fmt.Scanf("%s", &cambio2)
		fmt.Print("que deseas ponerle?")
		fmt.Scanf("%s", &valor2)
		p1.cambio(cambio2, valor2)
		p1.Print()

	} else {
		fmt.Println("Muchas gracias")
	}

}

type Persona struct {
	nombre, apellido, correo, contraseña, edad string
}

func (p *Persona) cambio(cambio2, valor2 string) {

	switch cambio2 {
	case nombre:
		p.cambiarNombre(valor2)
	case apellido:
		p.cambiarApellido(valor2)
	case correo:
		p.cambiarCorreo(valor2)
	case contraseña:
		p.cambiarContraseña(valor2)
	case edad:
		p.cambiarEdad(valor2)
	}
}

func (p *Persona) cambiarNombre(nombre string) string {
	p.nombre = nombre
	return p.nombre
}

func (p *Persona) cambiarApellido(apellido string) string {
	p.apellido = apellido
	return p.apellido
}

func (p *Persona) cambiarEdad(edad string) string {
	p.edad = edad
	return p.edad
}

func (p *Persona) cambiarCorreo(correo string) string {
	p.correo = correo
	return p.correo
}

func (p *Persona) cambiarContraseña(contraseña string) string {
	p.contraseña = contraseña
	return p.contraseña
}

func (p Persona) Print() {
	fmt.Printf("tus datos son: \nnombre: %v\napellido:%v\ncorreo:%v\ncontraseña:%v\nedad:%v", p.nombre, p.apellido, p.correo, p.contraseña, p.edad)
}
