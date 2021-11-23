package ej4

import "fmt"

func Ej4() {
	var apellido string = "Gomez"
	var edad int = 35
	boolean := false
	var sueldo float64 = 45857.90
	var nombre string = "Julián"

	fmt.Printf("\nHola soy %s %s, tengo %d años, soy un %t y mi sueldo es de %.1f", nombre, apellido, edad, boolean, sueldo)
}
