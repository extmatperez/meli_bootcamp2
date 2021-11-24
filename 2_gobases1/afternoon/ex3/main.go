package main

import "fmt"

// Préstamo
func main() {

	cliente_edad := 25
	cliente_empleado := true
	cliente_antiguedad := 2
	cliente_sueldo := 150000

	if cliente_edad > 22 {
		if cliente_empleado {
			if cliente_antiguedad > 1 {
				if cliente_sueldo > 100000 {
					fmt.Println("Prestamo otorgado")		
				} else {
					fmt.Println("Sueldo insuficiente")		
				}
			} else {
				fmt.Println("El cliente no tiene antiguedad suficiente")		
			}
		} else {
			fmt.Println("El cliente no está empleado")	
		}
	} else {
		fmt.Println("Edad no permitida")
	}
	
}
