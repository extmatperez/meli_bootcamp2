package main

import "fmt"

func main() {

	salary1 := 40000.0
	salary2 := 55000.0
	salary3 := 160000.0

	fmt.Printf("Para el salario %.2f, el impuesto es %.2f\n", salary1, taxFor(salary1))
	fmt.Printf("Para el salario %.2f, el impuesto es %.2f\n", salary2, taxFor(salary2))
	fmt.Printf("Para el salario %.2f, el impuesto es %.2f\n", salary3, taxFor(salary3))
}

func taxFor(salary float64) float64 {

	if salary < 50000 {
		return 0
	}

	if salary < 150000 {
		return salary * 0.17
	}

	return salary * 0.27
}
