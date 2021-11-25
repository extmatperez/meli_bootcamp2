package main

import "fmt"

func main() {
	fmt.Println("El salario total es", salary(6000, "a"))
}

func salary(minutes float64, category string) float64 {
	var (
		hours float64
		salaryTotal float64
	)

	hours = minutes / 60

	switch category {
	case "c": salaryTotal = hours * 1000
	case "b": salaryTotal = (hours * 1500) + ((hours * 1500) * 0.20)
	case "a": salaryTotal = (hours * 3000) + ((hours * 3000) * 0.50)
	}

	return salaryTotal

}