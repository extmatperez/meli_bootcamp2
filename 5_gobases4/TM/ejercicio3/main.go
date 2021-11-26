package main

import (
	"fmt"
)

func main() {
	salary := 100000
	err := CalcularImpuesto(salary)
	if err != nil {
		fmt.Println(err)
	}

}

func CalcularImpuesto(salary int) error {

	if salary < 150000 {
		return fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salary)
	} else {
		fmt.Printf("Debe pagar impuesto\n")
		return nil
	}
}
