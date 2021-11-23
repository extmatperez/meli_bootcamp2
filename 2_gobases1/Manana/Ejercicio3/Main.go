package main

import "fmt"

func main()  {
	/*
	var 1nombre string
 	var apellido string
  	var int edad
  	1apellido := 6
  	var licencia_de_conducir = true
  	var estatura de la persona int
  	cantidadDeHijos := 2
	*/

	//No puede arrancar con un numero
	var nombre string
	var apellido string
	var edad int
	//Si se usan los := no es necesario definirlo antes
	apellido = "Pallua"
	//Se puede definir directamente usando :=
	licencia_de_conducir := true
	//Si se va a definir se tiene que usar
	//No pueden ser palabras sueltas
	var estatura_de_la_persona int
	estatura_de_la_persona = 175

	cantidadDeHijos := 2

	fmt.Println("Correccion")
	//Todas las variables definidas deben ser usadas
	nombre = "Patricio"
	apellido = "Pallua"
	edad = 23
	fmt.Println("Hola",nombre,apellido,"tenes",edad,"años")
	fmt.Println("Tu altura es",estatura_de_la_persona,"y por suerte no tenes",cantidadDeHijos,"hijos")
	fmt.Println(licencia_de_conducir)
}
