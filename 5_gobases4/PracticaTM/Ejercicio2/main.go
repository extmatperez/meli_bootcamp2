package main

import (
	"errors"
	"fmt"
)

func verificarError(salario int) error {
	if salario < 150000 {
		return errors.New("error: el salario ingresado no alcanza el mínimo imponible")
	}
	return nil
}

func main() {
	var salary = 100000
	err := verificarError(salary)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuesto")
	}

	salary2 := 160000
	err = verificarError(salary2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}
