package main

import (
	"fmt"
	"errors"
)

type MyError struct {
	salario int
	msg	   string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Salario %d - %s", e.salario, e.msg)
}

func Ej2(salary int) (string, error) {

	if salary < 150000 {
		return "", errors.New("error: el salario ingresado no alcanza el mínimo imponible")
	}
	return "debe pagar impuesto", nil
}

func main() {
	fmt.Println(Ej2(140000))
	fmt.Println(Ej2(160000))
}