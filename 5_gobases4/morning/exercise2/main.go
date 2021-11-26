package main

import (
	"errors"
	"fmt"
)

type myError struct {
	s string
}

func errorControl(salary int) (int, error) {

	if salary < 150000 {
		return 0, errors.New("error: el salario ingresado no alcanza el mínimo imponible")
	}
	return salary, nil
}

func main() {
	var salary int
	salary = 150001
	response, err := errorControl(salary)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El salario ingresado es: ", response)
	}
}
