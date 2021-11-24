package main

import "fmt"

// Calcular salario
const (
	A = "a"
	B = "b"
	C = "c"
)

func calcSalary(minutes int, category string) float64 {
	hour := minutes / 60
	final_salary := 0.0
	switch category {
		case A:
			final_salary = float64(1000 * hour)
			return final_salary
		case B:
			final_salary = (float64(1500 * hour)) * 1.2
			return final_salary
		case C:
			final_salary = (float64(3000 * hour)) * 1.5
			return final_salary
		default:
			return final_salary
	}
}

func main() {
	//fmt.Println(calcSalary(120.0, "a"))
	//fmt.Println(calcSalary(120.0, "b"))
	fmt.Println(calcSalary(120.0, "c"))
}