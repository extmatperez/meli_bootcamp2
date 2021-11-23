package main

import (
	"fmt"
)

func main() {
	/*var apellido string = "Gomez"
	var edad int = "35"
	boolean := "false";
	var sueldo string = 45857.90
	var nombre string = "Julián"*/

	var apellido string = "Gomez"
	var edad int = 35
	boolean := false
	var sueldo float32 = 45857.9
	var nombre string = "Julián"

	fmt.Printf("Apellido: %s \nEdad:%d\nBoolean=%v\nSueldo: %f\nNombre:%s", apellido, edad, boolean, sueldo, nombre)

}
