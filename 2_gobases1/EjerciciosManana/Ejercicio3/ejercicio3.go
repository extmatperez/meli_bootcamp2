package main

import "fmt"

func main() {

	//var 1nombre string -- incorrecta porque inicia con numero
	//correcta:
	var nombre1 string
	nombre1 = "nombre1"
	//variable correcta
	var apellido string
	apellido = "apellido"
	// var int edad  -- variable incorrecta nombre y tipo estan en orden cambiado
	//correcta
	var edad int
	edad = 24
	// 1apellido := 6 -- variable incorrecta no puede iniciar el nombre con un numero
	//correcta:
	apellido1 := 6

	var licencia_de_conducir = true

	// var estatura de la persona int -- variable incorrecta no puede contener espacios en el nombre
	//correcta:
	var estatura_de_la_persona int
	estatura_de_la_persona = 170
	cantidadDeHijos := 2

	fmt.Println("variables: \nnombre1 = ", nombre1, "\napellido = ", apellido, "\nedad = ", edad, "\napellido1 = ", apellido1, "\nlicencia_de_conducir = ", licencia_de_conducir, "\nestatura_de_la_persona = ", estatura_de_la_persona, "\ncantidadDeHijos = ", cantidadDeHijos)

}
