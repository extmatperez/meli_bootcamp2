package main

import (
	"fmt"
)

type myCustomError struct {
	msg string
}

func (e *myCustomError) Error() string {
	return fmt.Sprintf(e.msg)
}

func myCustomErrorTest(salary int) error {
	if salary < 150000 {
		return &myCustomError{
			msg: "error: el salario ingresado no alcanza el mínimo imponible",
		}
	} else {
		return &myCustomError{
			msg: "debe pagar impuesto",
		}
	}
}

func main() {

	var salary int

	fmt.Println("Ingrese un valor de salario")
	fmt.Scanln(&salary)

	err := myCustomErrorTest(salary)

	fmt.Println(err)
}
