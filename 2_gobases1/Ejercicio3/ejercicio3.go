package main

import "fmt"

func main() {

	/*
			var 1nombre string // incorrecta
		  	var apellido string // correcta
		  	var int edad //incorrecta
		  	1apellido := 6 //incorrecta
		  	var licencia_de_conducir = true //correcta
		  	var estatura de la persona int //incorrecta
		  	cantidadDeHijos := 2 //correcta
	*/

	var nombre1 string
	var apellido string
	var edad int
	apellido1 := 6
	var licencia_de_conducir = true
	var estatura_de_la_persona int
	cantidadDeHijos := 2

	fmt.Printf("%T\n", nombre1)
	fmt.Printf("%T\n", apellido)
	fmt.Printf("%T\n", edad)
	fmt.Printf("%T\n", apellido1)
	fmt.Printf("%T\n", licencia_de_conducir)
	fmt.Printf("%T\n", estatura_de_la_persona)
	fmt.Printf("%T\n", cantidadDeHijos)

}
