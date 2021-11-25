/* Ejercicio 1 - Impuestos de salario
Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo, para cumplir el objetivo es necesario
crear una función que devuelva el impuesto de un salario.
Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17% del sueldo y si gana más de $150.000 se le descontará además un 10%.
*/

package main

import "fmt"

func calculate(salary float64) {
	tax := 0.17
	tax2 := 0.27

	if salary > 50000.00 && salary <= 150000 {
		fmt.Printf("the tax value is: %.2f \n", salary*tax)
	} else if salary > 150000.00 {
		fmt.Printf("the tax value is: %.2f \n", salary*tax2)
	} else {
		fmt.Println("The salary is under $50000, you don't have taxes")
	}
}

func main() {
	var salary float64
	fmt.Printf("Enter the employee salary: ")
	fmt.Scanf("%f\n", &salary)
	calculate(salary)
}
