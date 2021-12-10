package main

import "fmt"

func main() {
	var salary int
	salary = 180

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

func (e *myError) Error() string {
	return e.s
}

func errorControl(salary int) (int, error) {
	if salary < 150000 {
		return 0, &myError{s: "error: el salario ingresado no alcanza el mínimo imponible"}
	}
	return salary, nil
}
