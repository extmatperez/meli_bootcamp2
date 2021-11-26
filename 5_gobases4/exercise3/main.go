package main

import (
	"fmt"
)

type myError struct {
	s string
}

func errorControl(salary int) (int, error) {

	if salary < 150000 {
		return 0, fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salary)
	}
	return salary, nil
}

func main() {
	var salary int
	salary = 15000
	response, err := errorControl(salary)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El salario ingresado es: ", response)
	}
}
