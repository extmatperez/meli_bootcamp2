package main

import (
	"errors"
	"fmt"
)

func pagaImpuesto(salary int) (string, error) {

	if salary < 150000 {
		return "", errors.New("error: el salario ingresado no alcanza el mínimo imponible")
	} else {
		return "Debe pagar impuesto", nil
	}

}

func main() {

	salary := 10

	respuesta, err := pagaImpuesto(salary)

	if err != nil {
		fmt.Printf("Ocurrio un error %v\n", err.Error())
	} else {
		fmt.Println(respuesta)
	}

}
