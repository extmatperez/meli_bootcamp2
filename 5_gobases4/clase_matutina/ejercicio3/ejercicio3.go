package main

import "fmt"

func main() {
	salary := 3000

	msg, err := customErrorTest(salary)
	if err == nil {
		fmt.Printf("%v\n", msg)
	} else {
		fmt.Printf("%v\n", err)
	}
}

type CustomError struct {
	msg string
}

func customErrorTest(salary int) (string, error) {
	if salary < 150000 {
		return "", fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %v", salary)
	} else {
		return "Debe pagar impuesto", nil
	}
}

func (e CustomError) Error() string {
	return fmt.Sprintf("%v\n", e.msg)
}
