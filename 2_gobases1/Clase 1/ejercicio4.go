package main

import "fmt"

func main() {
	var apellido string = "Gomez"
	// var edad int = "35"
	var edad int = 35
	boolean := "false" //estaria mal asignar un string a un booleano pero como solo es una etiqueta esta bien
	// var sueldo string = 45857.90
	var sueldo float32 = 45857.90
	var nombre string = "Julián"

	fmt.Printf(apellido, edad, boolean, sueldo, nombre)
}
