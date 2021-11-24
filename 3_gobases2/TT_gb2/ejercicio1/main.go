package main

import "fmt"

// Ejercicio 1 - Impuestos de Salario
// Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo, para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario. 
// Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17% del sueldo y si gana más de $150.000 se le descontará además un 10%.

func tax_calculation(income float64) float64 {
	tax1 := 0.17
	tax2 := 0.27

	if income > 50000.00 {
		return income - income*tax1
	} else if income > 150000.00 {
		return income - income*tax2
	} else {
		return income
	}
}

func main () {

	var employee_income float64
	
	fmt.Printf("Enter employee income: ")
	fmt.Scanf("%f\n", &employee_income)

	fmt.Printf("Final income value: %.2f\n", tax_calculation(employee_income))


}