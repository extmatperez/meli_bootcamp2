package main

import (
	"errors"
	"fmt"
)

func Ej2(salary int) (string, error) {
	if salary < 150000 {
		return "", errors.New("error: el salario ingresado no alcanza el mínimo imponible")
	}
	return "Debe pagar impuestos", nil
}

func main() {
	fmt.Println(Ej2(1000))
}
