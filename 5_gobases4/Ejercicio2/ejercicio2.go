package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println()
	salary := 160.000
	if salary <= 150.000 { // si err no es nulos es por que existe un error
		fmt.Println(errors.New("error: el salario ingresado no alcanza el minimo imponible"))
	} else {
		fmt.Println("Debe pagar impuestos")
	}
	fmt.Println()
}
