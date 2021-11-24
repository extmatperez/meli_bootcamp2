package main

import "fmt"

func main() {

	const RECEIVE_LOAN = "Esta persona puede recibir prestamo con interes"
	const NOT_RECEIVE_LOAN = "Esta persona no puede recibir prestamo"
	const RECEIVE_LOAN_WITHOUT_INTEREST = "Esta persona puede recibir prestamo sin interes"

	var age int
	var employeer string
	var yearsOfWorked int
	var baseSalary float32
	fmt.Println("Please enter your age: ")
	fmt.Scanf("%d\n", &age)
	fmt.Println("Are you working now? enter \"yes\" or \"no\": ")
	fmt.Scanf("%s\n", &employeer)
	fmt.Println("Years of working: ")
	fmt.Scanf("%d\n", &yearsOfWorked)
	fmt.Println("Base salary: ")
	fmt.Scanf("%f.2\n", &baseSalary)
	fmt.Println(baseSalary)

	var isEmployeer bool
	switch employeer {
	case "yes":
		isEmployeer = true
	case "no":
		isEmployeer = false
	default:
		isEmployeer = false
	}

	if age >= 22 && isEmployeer && yearsOfWorked > 1 {
		if baseSalary > 100000.00 {
			fmt.Println(RECEIVE_LOAN_WITHOUT_INTEREST)
		} else {
			fmt.Println(RECEIVE_LOAN)
		}
	} else {
		fmt.Println(NOT_RECEIVE_LOAN)
	}
}
