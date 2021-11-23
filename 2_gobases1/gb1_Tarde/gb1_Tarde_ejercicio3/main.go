package main

import "fmt"

// Ejercicio 3 - Préstamo
// Un Banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos. Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar. Solo le otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y tengan más de un año de antigüedad en su trabajo. Dentro de los préstamos que otorga no les cobrará interés a los que su sueldo es mejor a $100.000. 
// Es necesario realizar una aplicación que tenga estas variables y que imprima un mensaje de acuerdo a cada caso.
// Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.

func main () {
	
	age := 24
	working := true
	years_working := 3
	income := 100000

	if age > 22 && working == true && years_working > 1 {
		if income > 100000 {
			fmt.Printf("El cliente cumple las condiciones para el otorgamiento del préstamo. SIN INTERÉS.")
		} else {
			fmt.Printf("El cliente cumple las condiciones para el otorgamiento del préstamo. CON INTERÉS.")
		}

	} else {
		fmt.Printf("El cliente no cumple con las condiciones para el otorgamiento del préstamo.")
	}


}