package main

import (
	"fmt"
)

func main() {
	var salary int
	fmt.Println("Ingrese su salario:")
	fmt.Scanf("%d", &salary)
	should_pay_taxes := tax_calculator(salary)
	if should_pay_taxes != nil {
		fmt.Printf("%v\n", should_pay_taxes)
	} else {
		fmt.Println("Debe pagar impuesto")
	}

}

func tax_calculator(salary int) error {
	if salary > 150000 {
		return nil
	} else {
		return fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salary)
	}
}
