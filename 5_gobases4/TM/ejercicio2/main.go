package main

import (
	"errors"
	"fmt"
)

func main() {
	salary := 200000
	err := CalcularImpuesto(salary)
	if err != nil {
		fmt.Println(err)
	}

}

func CalcularImpuesto(salary int) error {

	if salary < 150000 {
		return errors.New("error: el salario ingresado no alcanza el mínimo imponible")
	} else {
		fmt.Printf("Debe pagar impuesto\n")
		return nil
	}
}
