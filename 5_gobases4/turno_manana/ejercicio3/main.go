// Ejercicio 3 - Impuestos de salario #3
// Repite el proceso anterior, pero ahora implementando “fmt.Errorf()”, para que el mensaje de error reciba por parámetro el valor de “salary” indicando que no alcanza el mínimo imponible (el mensaje mostrado por consola deberá decir: “error: el mínimo imponible es de 150.000 y el salario ingresado es de: [salary]”, siendo [salary] el valor de tipo int pasado por parámetro).

package main

import (
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
	minimo := 150000
	if salary < minimo {
		return "", fmt.Errorf("\nEl mínimo imponible es de %d y el salario ingresado es de: %d\n", minimo, salary)
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
