package main

import (
	"fmt"
)

func main() {
	salary := 500

	if salary < 15000 {
		err := fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %v", salary)
		fmt.Println(err)
	} else {
		fmt.Println("Correcto")

	}

}
