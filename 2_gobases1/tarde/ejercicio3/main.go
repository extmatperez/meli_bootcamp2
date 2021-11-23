package main

import "fmt"

func main() {
	/*
		Un Banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos. Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar. Solo le otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y tengan más de un año de antigüedad en su trabajo. Dentro de los préstamos que otorga no les cobrará interés a los que su sueldo es mejor a $100.000.
		Es necesario realizar una aplicación que tenga estas variables y que imprima un mensaje de acuerdo a cada caso.
	*/

	edad := 23
	empleado := true
	antiguedad := 3
	sueldo := 1000

	if edad > 22 && empleado && antiguedad > 1 {
		if sueldo > 100000 {
			fmt.Println("Se le otorga prestamo con interes")
		} else {
			fmt.Println("Se le otorga prestamo sin interes")
		}
	} else {
		fmt.Println("No se le otorga prestamo")
	}

}
