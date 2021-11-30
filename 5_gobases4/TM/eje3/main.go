package main

import (
	"fmt"
	"os"
)

type myCustomError struct {
	msg string
}

func (e *myCustomError) Error() string {
	return fmt.Sprintf("%v", e.msg)
}

func myCustomErrorTest(salary int) error {
	if salary < 150000 {
		return &myCustomError{
			msg: fmt.Sprintln(fmt.Errorf("error: el minimo imponible es de 150.000 y el salario ingresado es de %v", salary)),
		}
	}
	return nil
}

func main() {

	salary := 145000
	err := myCustomErrorTest(salary)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%v No debe pagar impuestos", salary)
}
