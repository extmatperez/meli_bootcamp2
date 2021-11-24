package main

import "fmt"

func main() {

	fmt.Println(salary(60, "A"))
	fmt.Println(salary(60, "B"))
	fmt.Println(salary(60, "C"))
}

func salary(minutesWorked int, category string) float64 {

	hoursWorked := float64(minutesWorked) / 60.0

	switch category {

	case "A":
		return hoursWorked * 3000 * 1.5
	case "B":
		return hoursWorked * 1500 * 1.2
	default:
		return hoursWorked * 1000
	}
}
