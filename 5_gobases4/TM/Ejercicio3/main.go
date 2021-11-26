package main

import (
	"fmt"
)

func funcError(salary int) error {
	if salary < 150000 {
		return fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salary)
	}
	return nil
}

func main() {
	salary := 150
	err := funcError(salary)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}
