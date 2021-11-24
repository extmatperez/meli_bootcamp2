package main

import "fmt"

// Impuestos de salario
func calcTax(salary float64) float64 {
	tax := 0.0
	if salary > 50000 {
		tax = salary * 0.17
		if salary > 150000 {
			tax = salary * 0.27
		}
	}
	return tax
}

func main() {

	salary_1 := 160000.0
	salary_2 := 110000.0
	final_tax_1 := calcTax(salary_1)
	final_tax_2 := calcTax(salary_2)

	fmt.Println("Impuesto final sobre salario:", final_tax_1)
	fmt.Println("Impuesto final sobre salario:", final_tax_2)

}