package main

import "fmt"

func main()  {
	
	/*
		var apellido string = "Gomez"
		var edad int = "35"
		boolean := "false";
		var sueldo string = 45857.90
		var nombre string = "Julián"
	*/

	var apellido string = "Gomez"
	fmt.Println("Hola",apellido)
	//Si se define como "int" no puede ser un string
	var edad int = 35
	fmt.Println("Tu edad es",edad,"años")
	//Esta tecnicamente bien porque boolean no es palabra reservada
	//Entonces boolean es una variable que guarda un string
	//Y el ";" al final tampoco esta mal, no es necesario pero no rompe
	boolean := false 
	fmt.Println(boolean)
	//No se puede definir como string y pasarle un numero float
	var sueldo float32 = 45857.90
	fmt.Println("Tu sueldo es",sueldo)
	//En este no encontre error
	var nombre string = "Julian"
	fmt.Println("Mucho gusto",nombre)
}