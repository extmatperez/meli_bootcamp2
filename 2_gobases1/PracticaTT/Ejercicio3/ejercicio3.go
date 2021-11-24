package main

import "fmt"

func main() {
	edad := 24
	estado := "Empleado"
	antiguedad := 3
	sueldo := 88000

	if edad > 22 {
		if estado == "Empleado" || estado == "empleado" {
			if antiguedad > 1 {
				if sueldo > 100000 {
					fmt.Println("El cliente puede acceder a un préstamo sin interés.")
				} else {
					fmt.Println("El cliente puede acceder a un préstamo con interés.")
				}
			} else {
				fmt.Println("El cliente posee una antiguedad menor a 1 año, no puede acceder al préstamo.")
			}
		} else {
			fmt.Println("El cliente se encuentra desempleado, no puede acceder al préstamo.")
		}
	} else {
		fmt.Println("El cliente no es mayor de 22 años, no puede acceder al préstamo.")
	}
}
