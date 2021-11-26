package main

import (
	"fmt"
)

func verificarError(salario int) error {
	if salario < 150000 {
		return fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salario)
	}
	return nil
}

func main() {
	var salary = 100000
	err := verificarError(salary)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuesto")
	}

	salary2 := 160000
	err = verificarError(salary2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}
