package main

import (
	"fmt"
)

func main() {
	var salary int

	fmt.Println("Ingrese un valor de salario")
	fmt.Scanln(&salary)

	err := fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salary)

	if salary < 150000 {

		fmt.Println(err)
	} else {
		fmt.Println("debe pagar impuesto")
	}

}
