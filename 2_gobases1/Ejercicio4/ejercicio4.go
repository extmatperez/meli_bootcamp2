package main

import "fmt"

func main() {

	/*
		var apellido string = "Gomez" correcta
		var edad int = "35" incorreta
		boolean := "false"; correcta
		var sueldo string = 45857.90 incorrecta
		var nombre string = "Julián" correcta
	*/

	var apellido string = "Gomez"
	var edad int = 35
	boolean := "false"
	var sueldo string = "45857.90"
	var nombre string = "Julián"

	fmt.Printf("%T\n", apellido)
	fmt.Printf("%T\n", edad)
	fmt.Printf("%T\n", boolean)
	fmt.Printf("%T\n", sueldo)
	fmt.Printf("%T\n", nombre)

}
