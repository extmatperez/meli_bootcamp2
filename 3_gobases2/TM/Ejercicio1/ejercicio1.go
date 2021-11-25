package main

import "fmt"

const (
	first_threshold  = 50000.00
	second_threshold = 1500000.00
)

func main() {
	salary := 0.00
	fmt.Println("Ingrese sue sueldo: ")
	fmt.Scanf("%f", &salary)
	fmt.Printf("Sus impuestos son: %.2f\n", tax_calculator(salary))
}

func tax_calculator(salary float64) float64 {
	var tax float64
	switch salary > 0.0 {
	case first_threshold < salary && salary < second_threshold:
		tax = salary * 0.17
	case second_threshold < salary:
		tax = salary * 0.1
	default:
	}
	return tax
}
