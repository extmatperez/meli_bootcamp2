package main

import "fmt"

func main() {
	//Las variables correctas son: apellido (la primera vez), y cantidadDeHijos
	//Las variables mal declaradas son las siguientes:
	//1nombre empieza con un numero
	//edad está mal declarada porque tiene primero el tipo de dato y despues la variable
	//1apellido empieza con un número. Entiendo que es una variable diferente al string "Apellido" por lo que le cambio el nombre a "apellidoNum"
	//licencia_de_conducir no tiene el tipo de variable escrito "bool"
	//estatura de la persona tiene espacios, y las variables no admiten espacios
	var nombre string
	var apellido string
	var edad int
	apellidoNum := 6
	var licencia_de_conducir bool = true
	var estatura_de_la_persona int
	cantidadDeHijos := 2

}

