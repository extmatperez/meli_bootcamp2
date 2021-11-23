package main

import "fmt"

// Ejercicio 4
// Un estudiante de programación intentó realizar declaraciones de variables con sus respectivos tipos en Go pero tuvo varias dudas mientras lo hacía. A partir de esto, nos brindó su código y pidió la ayuda de un desarrollador experimentado que pueda:
// Verificar su código y realizar las correcciones necesarias.

//   var apellido string = "Gomez"		---> OK
//   var edad int = "35"				---> var edad int = 35
//   boolean := "false";				---> boolean := false
//   var sueldo string = 45857.90		---> var sueldo float = 45857.90
//   var nombre string = "Julián"		---> OK

func main () {

	var apellido string = "Gomez"
	var edad int = 35
	boolean := false
	var sueldo float64 = 45857.90
	var nombre string = "Julián"

	fmt.Printf("Estas son las declaraciones correctas: \n string %s \n integer %d \n booleano %T \n flotante %.2f \n string %s", apellido, edad, boolean, sueldo, nombre)

}


