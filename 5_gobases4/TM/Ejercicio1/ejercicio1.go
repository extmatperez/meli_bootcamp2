package main

import "fmt"

func main() {
	var salary int = 150000

	if salary < 150000 {
		fmt.Println(testMySalaryTaxError(salary))
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}

func testMySalaryTaxError(salary int) error {
	return &MySalaryTaxError{}
}

type MySalaryTaxError struct {
}

func (e *MySalaryTaxError) Error() string {
	return "error: el salario ingresado no alcanza el mínimo imponible"
}
