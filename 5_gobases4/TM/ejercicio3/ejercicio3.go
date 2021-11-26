package main

import (
	"fmt"
)

func Ej3(salary int) (string, error) {
	if salary < 150000 {
		return "", fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salary)
	}
	return "Debe pagar impuestos", nil
}

func main() {
	fmt.Println(Ej3(1000))
}
