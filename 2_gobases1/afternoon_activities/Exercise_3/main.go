/* Ejercicio 3 - Préstamo

Un Banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos.
Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar.
Solo le otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y tengan más de un año de antigüedad en su trabajo.
Dentro de los préstamos que otorga no les cobrará interés a los que su sueldo es mejor a $100.000.
Es necesario realizar una aplicación que tenga estas variables y que imprima un mensaje de acuerdo a cada caso. */

package main

import "fmt"

var age int
var employee bool
var antiquity float64
var salary float64

func main() {

	age = 23
	employee = true
	antiquity = 1.5
	salary = 143000

	if age > 22 {
		if employee == true {
			if antiquity > 1 {
				if salary > 100000 {
					fmt.Println("You don't have to pay interests!")
				} else {
					fmt.Println("You have to pay interests.")
				}
			} else {
				fmt.Println("You have to have more than 1 year like employee.")
			}
		} else {
			fmt.Println("You have to be employee.")
		}
	} else {
		fmt.Println("You have to be older than 22.")
	}

}
