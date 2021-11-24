package main

import "fmt"

func main() {

	var edad int
	var sueldo int
	var empleado bool
	var meses_en_el_mismo_empleo int

	edad = 23
	sueldo = 500000
	empleado = true
	meses_en_el_mismo_empleo = 13

	if edad <= 22 || !empleado || meses_en_el_mismo_empleo <= 12 {
		fmt.Println("No puede aplicar al prestamo")
	} else {

		if sueldo > 100000 {
			fmt.Println("Puede aplicar al prestamo y no se le cobrara interes")
		} else {
			fmt.Println("Puede aplicar al prestamo")
		}
	}

}
