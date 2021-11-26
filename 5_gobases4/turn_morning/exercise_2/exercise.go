package main

import (
	"errors"
	"fmt"
)

type myError struct {
	message string
	code    int
}

func (e *myError) Error() string {
	return e.message
}

func errorControl(salary int) (int, error) {

	if salary < 150000 {
		return 0, &myError{
			message: "error: el salario ingresado no alcanza el mínimo imponible",
			code:    400,
		}
	}
	return salary, nil
}

func main() {
	var salary int = 15000
	response, err := errorControl(salary)

	if err != nil {
		fmt.Println(errors.New(err.Error()))
	} else {
		fmt.Println("El salario ingresado es: ", response)
	}
}
