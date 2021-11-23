package main

import "fmt"

func main() {

	//var 1nombre string -- Los nombres de las variables deben comenzar con letras no con numero
	var nombre1 string
	nombre1 = "andres"
	fmt.Println(nombre1)

	var apellido string //correcta
	apellido = "pachon"
	fmt.Println(apellido)

	//var int edad -- Esta mal ordenada la forma de la declaracion  de la variable
	var edad int
	edad = 25
	fmt.Println(edad)

	//1apellido := 1 -- Los nombres de las variables no deben comenzar en letras ademas se estaria asignando como tipo de variable un entero cuando la variable hace referencia a un apellido
	apellido1 := "Pachón"
	fmt.Println(apellido1)

	//var licencia_de_conduccion = true -- No se a definido el tipo de la variable y si son varias palabras se empieza la segunda palabra con mayuscula y todo pegado
	licenciaDeConduccion := false
	fmt.Println(licenciaDeConduccion)

	//var estatura de la persona int -- El nombre de las variables no pueden ir con espacios debe ir todo junto y apartir de la segunda palabra debe empezar en mayuscula
	var estaturaDeLaPersona int
	estaturaDeLaPersona = 160
	fmt.Println(estaturaDeLaPersona)

	cantidadDeHijos := 2 //correcta
	fmt.Println(cantidadDeHijos)
}
