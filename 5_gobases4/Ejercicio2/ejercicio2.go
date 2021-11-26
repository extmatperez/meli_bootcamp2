package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println()
	salary := 160.000
	if salary <= 150.000 {
		fmt.Println(errors.New("error: el salario ingresado no alcanza el minimo imponible"))
	} else {
		fmt.Println("Debe pagar impuestos")
	}
	fmt.Println()
}
