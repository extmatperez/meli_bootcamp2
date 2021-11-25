package main

import "fmt"

func main(){
	var (
		initialSalary float64 = 170000
		mainTax, plusTax, salary = tax(initialSalary)
	)
	fmt.Println("El salario inicial fue de: ",initialSalary,
			"\nEl salario final es: ", salary,
			"\nEl impuesto aplicado fue de: ", mainTax,
			"\nY el impuesto agregado es: ", plusTax)
}

func tax(salary float64) (float64, float64, float64) {
	// declare variables
	var (
		mainTax float64
		plusTax float64
		finalSalary float64
	)
	// logic to calculate taxes
	if salary > 50000{
		mainTax = (salary * 0.15)
		finalSalary = salary - mainTax
		if salary > 150000{
			plusTax = (salary - mainTax) * 0.10
			finalSalary = finalSalary - plusTax
		}
	}
	// returning taxes and finalSalary
	return mainTax, plusTax, finalSalary
}