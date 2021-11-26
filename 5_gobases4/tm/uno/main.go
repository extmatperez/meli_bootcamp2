package main

import (
	"fmt"
	"os"
)

type customError struct {
	msg    string
	status int
}

func (e *customError) Error() string {
	return fmt.Sprintf("%d - %v", e.status, e.msg)
}

func testCustomError(salary int) (int, error) {
	if salary > 150000 {
		return 400, &customError{
			msg:    "error: el salario ingresado no alcanza el mínimo imponible",
			status: 400,
		}
	}
	return 200, nil
}

func main() {
	salary := 140000
	_, err := testCustomError(salary)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println("Debe pagar impuesto.")
	}
}
