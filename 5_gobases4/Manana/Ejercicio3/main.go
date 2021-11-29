package main

import (
	"fmt"
)

func pagaImpuesto(salary int) (string, error) {

	if salary < 150000 {
		return "", fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d\n", salary)
	} else {
		return "Debe pagar impuesto", nil
	}

}

func main() {

	salary := 150000

	respuesta, err := pagaImpuesto(salary)

	if err != nil {
		fmt.Printf("Ocurrio un error %v\n", err.Error())
	} else {
		fmt.Println(respuesta)
	}

}
