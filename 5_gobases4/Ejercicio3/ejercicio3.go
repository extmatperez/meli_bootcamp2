package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	salary := 10.000
	if salary <= 150.000 {
		err := fmt.Errorf("error: el minimo imponible es de 150.000 y el salario ingresado es de:%.3f ", salary)
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuestos")
	}
	fmt.Println()
}
