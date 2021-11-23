package main

import "fmt"

func main() {

	//var 1nombre string     Incorrecta, las var no comienzan con numeros
	//var apellido string    Correcta
	//var int edad           Incorrecta, primero se declara el nombre y luego el tipo
	//1apellido := 6		 Incorrecta, las var no comienzan con numeros y ya esta declarada como string
	//var licencia_de_conducir = true  Incorrecta, no especifica tipo de dato
	//var estatura de la persona int   Incorrecta, debe ser camelCase o snake_case
	//cantidadDeHijos := 2   Correcta

	var nombre string = "Nico"
	var apellido string = "Arguello"
	var edad int = 36
	//apellido := 6
	var licencia_de_conducir bool = true
	var estatura_de_la_persona float64 = 1.8
	cantidadDeHijos := 2

	fmt.Println(nombre)
	fmt.Println(apellido)
	fmt.Println(edad)
	fmt.Println(licencia_de_conducir)
	fmt.Println(estatura_de_la_persona)
	fmt.Println(cantidadDeHijos)

}
