package main

import "fmt"

func main() {

	// var 1nombre string => Incorrecta 
	var apellido string // => Correcta
	// var int edad // => Incorrecta
	// 1apellido := 6 => Incorrecta 
	var licencia_de_conducir = true
	// var estatura de la persona int => Incorrecta 
	cantidadDeHijos := 2
	
	/* CORRECIONES */
	var edad int
	var nombre string
	apellido_1 := 6
	var estatura_de_la_persona int

	fmt.Println(apellido, licencia_de_conducir, cantidadDeHijos, edad, nombre, apellido_1, estatura_de_la_persona)
}