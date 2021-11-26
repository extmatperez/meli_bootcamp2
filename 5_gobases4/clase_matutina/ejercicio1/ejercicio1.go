package main

import "fmt"

func main() {
	salary := 300000

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
		return "", &CustomError{
			msg: "error: el salario ingresado no alcanza el mínimo imponible",
		}
	} else {
		return "Debe pagar impuesto", nil
	}
}

func (e CustomError) Error() string {
	return fmt.Sprintf("%v\n", e.msg)
}
