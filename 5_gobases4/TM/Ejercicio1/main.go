package main

import "fmt"

type myError struct {
	msg string
}

func (e *myError) Error() string {
	(*e).msg = "error: el salario ingresado no alcanza el mínimo imponible"
	return (*e).msg
}

func funcError(salary int) error {
	if salary < 150000 {
		return (&myError{"error: el salario ingresado no alcanza el mínimo imponible"})
	}
	return nil
}

func main() {
	salary := 150000
	err := funcError(salary)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}
