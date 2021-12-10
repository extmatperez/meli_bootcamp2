package main

import (
	"errors"
	"fmt"
)

func main() {
	var salary int
	salary = 18000908

	response, err := errorControl(salary)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El salario es: ", response)
	}

}

type myError struct {
	s string
}

//func (e *myError) Error() string {
//	return e.s
//}

func errorControl(salary int) (int, error) {
	if salary < 150000 {
		return 0, errors.New("error: el salario ingresado no alcanza el mínimo imponible")
	}
	return salary, nil
}
