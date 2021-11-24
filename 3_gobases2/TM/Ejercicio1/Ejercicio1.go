package main

import "fmt"

func calculate_taxes(salary float64) float64 {
	var tax float64
	if salary > 50000 && salary <= 150000 {
		tax = salary * 0.17
	} else {
		if salary > 150000 {
			tax = salary * 0.27
		} else {
			tax = 0
		}
	}
	return tax
}

func main() {
	var salary float64
	fmt.Println("Salario: ")
	fmt.Scanf("%f", &salary)
	fmt.Printf("Los impuestos correspondientes son de %6.2f", calculate_taxes(salary))
}
