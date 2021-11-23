package main

import "fmt"

func main() {
	var edad = 25
	var empleado = true
	var antiguedad = 2
	var sueldo = 45000

	if edad >= 18 && empleado == true && antiguedad > 1 {
		if sueldo > 100000 {
			fmt.Printf("Puedes obtener un préstamo sin intereses\n")
		} else {
			fmt.Printf("Puedes obtener un préstamo con intereses\n")
		}
	} else {
		fmt.Printf("No puedes obtener un préstamo\n")
	}
}