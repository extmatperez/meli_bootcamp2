// Tipos de datos

/*
Un estudiante de programación intentó realizar declaraciones de variables con sus respectivos tipos en Go pero tuvo varias dudas mientras lo hacía.
A partir de esto, nos brindó su código y pidió la ayuda de un desarrollador experimentado que pueda:
Verificar su código y realizar las correcciones necesarias.

	var apellido string = "Gomez"
	var edad int = "35"
	boolean := "false";
	var sueldo string = 45857.90
	var nombre string = "Julián"



*/
package main

import "fmt"

func main() {
	var apellido string = "Gomez"
	var edad int = 35
	boolean := false
	var sueldo float64 = 45857.90
	var nombre string = "Julián"

	fmt.Printf("nombre completo: %s %s, edad: %d, licencia de conducir? %v, sueldo: %f", nombre, apellido, edad, boolean, sueldo)
}
