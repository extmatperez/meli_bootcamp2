/*
Haz lo mismo que en el ejercicio anterior pero reformulando el código para que, en reemplazo
de “Error()”, se implemente “errors.New()”.
*/

package main

import (
	"errors"
	"fmt"
)

type myError struct {
	msg string
}

func myErrorTest(salary int) (string, error) {

	if salary < 150000 {
		return "", errors.New("error: el salario ingresado no alcanza el mínimo imponible")
	} else {
		return "Debe pagar impuestos", nil
	}
}

func main() {

	salary := 16000

	msg, err := myErrorTest(salary)

	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Printf("%v\n", msg)
	}
}
