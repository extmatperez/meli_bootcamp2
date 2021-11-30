package main

import (
	"errors"
	"fmt"
	"os"
)

type myCustomError struct {
	msg string
}

func (e *myCustomError) Error() string {
	return fmt.Sprintf("%v", e.msg)
}

func myCustomErrorTest(salary int) error {
	if salary < 150000 {
		return &myCustomError{
			msg: fmt.Sprintln(errors.New("error: el salario no alcanza el mínimo imponible")),
		}

	}
	return nil
}

func main() {

	salary := 255000
	err := myCustomErrorTest(salary)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%v No debe pagar impuestos", salary)
}
