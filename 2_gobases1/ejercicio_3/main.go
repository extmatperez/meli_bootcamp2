package main

import "fmt"

func main() {

	//1nombre" -> "nombre"
	var nombre string

	var apellido string

	//"var int edad" -> "var edad int"
	var edad int

	//"apellido" -> "apellido1"
	//notar que es una variable diferente a "apellido", con un tipo distinto
	apellido1 := 6
	var licencia_de_conducir = true

	//"estatura de la persona" -> "estatura_de_la_persona"
	var estatura_de_la_persona int

	cantidadDeHijos := 2

	//agrego println para usar las variables y que el compilador no chille
	fmt.Println(nombre, apellido, apellido1, edad, licencia_de_conducir, estatura_de_la_persona, cantidadDeHijos)
}
