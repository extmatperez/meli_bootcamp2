package main

import "fmt"

func main() {

	var edad int
	var empleado bool
	var antiguedad int
	var sueldo float64
	var msn = "No se puede realizar prestamo por que: \n"

	edad = 22
	empleado = true
	antiguedad = 1
	sueldo = 101.000

	prestamoValido := true

	if edad < 22 {
		msn += "\tSu edad es menor a 22\n"
		prestamoValido = false
	}
	if !empleado {
		msn += "\tNo esta emepleado\n"
		prestamoValido = false
	}
	if antiguedad < 1 {
		msn += "\tTiene menos de un año laborando\n"
		prestamoValido = false
	}
	fmt.Println()
	if prestamoValido {
		fmt.Println("Ningun impedimento para realizar el prestamo")
	} else {
		fmt.Print(msn)
	}
	if sueldo > 100.000 {
		fmt.Println("\tSe cobraran intereses puesto que el salario es mayor a $100.000")
	}
}
