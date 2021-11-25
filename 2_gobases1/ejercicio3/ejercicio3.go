package main

import . "fmt"

func main() {
	/*var 1nombre string // incorrecta porque empieza con un número
	var apellido string //correcta
	var int edad // incorrecta. El tipo debe ir después del nombre
	1apellido := 6 //incorrecta porque empieza con número. Si es "apellido", es incorrecto porque cambia de tipo
	var licencia_de_conducir = true // correcto
	var estatura de la persona int //incorrecto. El nombre no puede tener espacios
	cantidadDeHijos := 2 // correcto
	*/
	var nombre string
	var apellido string
	var edad int
	apellido1 := 6
	var licencia_de_conducir = true //tipo bool
	var estatura_de_la_persona int
	cantidadDeHijos := 2
	nombre = "a"
	apellido = "b"
	estatura_de_la_persona = 0

	Println(nombre, apellido, edad, apellido1, licencia_de_conducir, estatura_de_la_persona, cantidadDeHijos)
}
