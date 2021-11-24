package main

import "fmt"

func main() {
	var salary float64 = 150000

	fmt.Printf("El salario con impuestos es: %.2f\n", calculateSalaryTax(salary))
}

func calculateSalaryTax(salary float64) float64 {
	salaryWithTax := salary
	if salary >= 50000 {
		salaryWithTax = salaryWithTax * 0.87

		if salary >= 150000 {
			salaryWithTax = salaryWithTax * 0.9
		}
	}

	return salaryWithTax
}
