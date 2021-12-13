package main

import "fmt"

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
		return &ErrorPersonalizado{}
	}
}

type ErrorPersonalizado struct {
}

func (e *ErrorPersonalizado) Error() string {
	return "Error: el salario ingresado no alcanza el mínimo imponible\n"
}
