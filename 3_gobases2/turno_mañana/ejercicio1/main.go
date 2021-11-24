// Ejercicio 1 - Impuestos de salario
// Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo, para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario.
// Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17% del sueldo y si gana más de $150.000 se le descontará además un 10%.

package main

import "fmt"

func calculate_discount(porc int, salari float32) float32 {
	if porc > 0 && porc <= 100 {
		return salari * float32(porc) / 100
	} else {
		return 0
	}
}

func calculate_tax(salari int) string {
	if salari > 50000 {
		discount := calculate_discount(17, float32(salari))
		discount_tostring := fmt.Sprintf("%f", discount)
		return discount_tostring

	} else if salari > 150000 {
		discount := calculate_discount(10, float32(salari))
		discount_tostring := fmt.Sprintf("%f", discount)
		return discount_tostring
	} else {
		return fmt.Sprintln("No es posible calcular el descuento")
	}
}

func main() {
	fmt.Println(calculate_tax(153000))
	fmt.Println(calculate_tax(100000))
}
