package main

import (
	"errors"
	"fmt"
)

func funcError(salary int) error {
	if salary < 150000 {
		return errors.New("error: el salario ingresado no alcanza el mínimo imponible")
	}
	return nil
}

func main() {
	salary := 1500
	err := funcError(salary)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}
