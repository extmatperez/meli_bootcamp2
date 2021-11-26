package main

import (
	"fmt"
	"os"
)

func errorTest(salary int) error {
	if salary < 150000 {
		return fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %v", salary)
	}
	return nil
}

func main() {
	var salary = 140000
	err := errorTest(salary)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Debe pagar impuesto.")
}
