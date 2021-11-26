package main

import (
	"fmt"
)

func crearError(salario int) (string, error) {
	if salario < 150000 {
		return "", fmt.Errorf("error: el salario ingresado no alcanza el mínimo imponible : %d", salario)
	} else {
		return "Debe pagar impuesto", nil
	}
}

func main() {
	var salario int = 2000

	ok, err := crearError(salario)

	if err == nil {
		fmt.Println(ok)
	} else {
		fmt.Println(err)
	}
}
