/*Un Banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos. Para ello tiene ciertas reglas
para saber a qué cliente se le puede otorgar. Solo le otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren
empleados y tengan más de un año de antigüedad en su trabajo. Dentro de los préstamos que otorga no les cobrará interés a los
que su sueldo es mejor a $100.000.
Es necesario realizar una aplicación que tenga estas variables y que imprima un mensaje de acuerdo a cada caso.

Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.
*/

package main

import (
	"fmt"
)

func main() {
	cliente := "Juan"
	edad := 20
	empleado := true
	antiguedad := 2
	sueldo := 160000

	if edad >= 22 && empleado && antiguedad >= 1 {
		if sueldo > 100000 {
			fmt.Println(cliente, "puede recibir préstamo pero se le cobrarán intereses")
		} else {
			fmt.Println(cliente, "puede recibir préstamo sin intereses")
		}
	} else {
		fmt.Println(cliente, "no puede recibir préstamo")
	}

}
