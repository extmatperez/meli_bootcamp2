package main

import (
	"errors"
	"fmt"
)

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
		return "", errors.New("error: el salario ingresado no alcanza el mínimo imponible")
	} else {
		return "Debe pagar impuesto", nil
	}
}
