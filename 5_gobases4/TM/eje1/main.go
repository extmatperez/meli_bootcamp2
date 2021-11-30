package main

import (
	"fmt"
	"os"
)

type myCustomError struct {
	salary int
	msg    string
}

func (e *myCustomError) Error() string {
	return fmt.Sprintf("%d - %v", e.salary, e.msg)
}

func myCustomErrorTest(salary int) (int, error) {
	if salary < 150000 {
		return 400, &myCustomError{
			salary: salary,
			msg:    "error: el salario no alcanza el mínimo imponible",
		}
	}
	return 200, nil
}

func main() {

	salary := 5000
	status, err := myCustomErrorTest(salary)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Debe pagar impuestos", status)
}
