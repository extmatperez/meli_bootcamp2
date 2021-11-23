package main

import "fmt"

func main() {
	var edad int
	var estado_empleo string
	var antiguedad int
	var sueldo int

	fmt.Println("Ingrese su edad:")
	fmt.Scanln(&edad)
	fmt.Println("¿Usted es empleado? (y/n)")
	fmt.Scanln(&estado_empleo)
	fmt.Println("Ingrese años de antiguedad")
	fmt.Scanln(&antiguedad)

	if edad > 21 {
		if estado_empleo == "y" {
			if antiguedad > 0 {
				fmt.Println("¡Felicitaciones! Es elegible para optar por un prestamo, ingrese su sueldo")
				fmt.Scanln(&sueldo)
				if sueldo > 100000 {
					fmt.Println("Usted no deberá pagar intereses")
				} else {
					fmt.Println("Solicite cotizacion con intereses")
				}
			} else {
				fmt.Println("Usted no es aplicable para solicitar un prestamo, debe tener 1 o más años de antiguedad")
			}
		} else {
			fmt.Println("Usted no es aplicable para solicitar un prestamo, debe encontrarse empleado")
		}
	} else {
		fmt.Println("Usted no es aplicable para solicitar un prestamo, no tiene la edad suficiente (>21 años)")
	}

}
