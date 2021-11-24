/*Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo,
para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario.
Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17% del sueldo y
si gana más de $150.000 se le descontará además un 10%.*/

package main

import "fmt"

func impuestoSueldo(sueldo float64) float64 {

	impuesto := 0.00

	if sueldo > 50000.0 && sueldo < 150000.0 {
		impuesto = sueldo * 0.17
	} else {
		if sueldo > 150000.0 {
			impuesto = sueldo * 0.10
		}
	}

	return impuesto
}

func main() {

	sueldo := 0.0

	fmt.Println("Ingrese el Sueldo: ")
	fmt.Scanf("%f", &sueldo)

	fmt.Println("Los impuestos correspondientes son de ", impuestoSueldo(sueldo))

}
