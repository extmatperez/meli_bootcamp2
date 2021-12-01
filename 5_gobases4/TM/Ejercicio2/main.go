package main

import (
	"errors"
	"fmt"
)

func main() {

	var salary int

	fmt.Println("Ingrese un valor de salario")
	fmt.Scanln(&salary)

	if salary < 150000 {
		fmt.Println(errors.New("error: el salario ingresado no alcanza el mínimo imponible"))
	} else {
		fmt.Println(errors.New("debe pagar impuesto"))
	}
}
