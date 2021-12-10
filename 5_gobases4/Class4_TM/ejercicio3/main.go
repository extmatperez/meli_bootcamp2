package main

import (
	"fmt"
)

func main() {
	var salary int
	salary = 1800

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
		return 0, fmt.Errorf("error: el salario ingresado no alcanza el mínimo imponible. Su salario es: %d", salary)
	}
	return salary, nil
}
