package main

import (
	"errors"
	"fmt"
	"os"
)

func errorTest(salary int) (int, error) {
	if salary < 150000 {
		return 400, errors.New("error: el salario ingresado no alcanza el mínimo imponible")
	}
	return 200, nil
}

func main() {
	salary := 130000
	_, err := errorTest(salary)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println("Debe pagar impuesto.")
	}
}
