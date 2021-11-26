/*
? Ejercicio 2 - Impuestos de salario #2

Haz lo mismo que en el ejercicio anterior pero reformulando el código para que,
en reemplazo de “Error()”,  se implemente “errors.New()”.
*/

package main

import (
	"errors"
	"fmt"
)

func main() {
	var salary int = 100000
	if salary <= 150000 {
		fmt.Println(errors.New("el salario ingresado no alcanza el mínimo imponible"))
	} else {
		fmt.Println("Debe pagar impuestos")
	}

	salary = 200000
	if salary <= 150000 {
		fmt.Println(errors.New("el salario ingresado no alcanza el mínimo imponible"))
	} else {
		fmt.Println("Debe pagar impuestos")
	}
}
