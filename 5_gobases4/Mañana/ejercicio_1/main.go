package main

import (
	"fmt"
)

type myError struct {
	status int
	msg    string
}

func (e *myError) Error() string {
	return "error: el salario ingresado no alcanza el mínimo imponible"
}

func test(n int) error {
	if n < 15000 {
		return &myError{

			msg: "error: el salario ingresado no alcanza el mínimo imponible",
		}
	}
	return nil
}
func main() {
	salary := 5

	err := test(salary)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Correcto")

	}

}
