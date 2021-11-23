package main

import . "fmt"

func main() {
	/*
			  var apellido string = "Gomez" // correcto
		  var edad int = "35" //incorrecto. No puede ser string
		  boolean := "false"; // incorrecto. No puede ser string
		  var sueldo string = 45857.90 //incorrecto. Tiene que declararse como float
		  var nombre string = "Julián" //correcto

	*/

	var apellido string = "Gomez" // correcto
	var edad int = 35             //incorrecto. No puede ser string
	boolean := false              // incorrecto. No puede ser string
	var sueldo float32 = 45857.90 //incorrecto. Tiene que declararse como float
	var nombre string = "Julián"  //correcto
	Printf("%s, %d, %t, %f, %s", apellido, edad, boolean, sueldo, nombre)
}
