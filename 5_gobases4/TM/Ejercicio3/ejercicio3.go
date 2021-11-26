package main

import (
	"fmt"
)

func main() {
	var salary int = 10000

	if salary < 150000 {
		fmt.Println(testMySalaryTaxError(salary))
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}

func testMySalaryTaxError(salary int) error {
	return fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salary)
}
