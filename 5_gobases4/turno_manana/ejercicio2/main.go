// Ejercicio 2 - Impuestos de salario #2

// Haz lo mismo que en el ejercicio anterior pero reformulando el código para que, en reemplazo de “Error()”,  se implemente “errors.New()”.

package main

import (
	"errors"
	"fmt"
	"os"
)

type myCustomError struct {
	msg string
}

// func (e myCustomError) Error() string {
// 	return fmt.Sprintf("%v", e.msg)
// }

func myCustomErrorTest(salary int) (string, error) {
	if salary < 150000 {
		return "", errors.New("\nEl salario ingresado no alcanza el mínimo imponible \n")
	}
	return "\nDebe pagar impuesto\n", nil
}

func main() {
	var salary int
	salary = 140000
	result, err := myCustomErrorTest(salary)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result)
}
