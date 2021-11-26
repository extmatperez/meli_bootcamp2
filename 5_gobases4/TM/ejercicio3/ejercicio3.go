package main

/*
Repite el proceso anterior, pero ahora implementando “fmt.Errorf()”, para que el mensaje de error reciba por parámetro el valor de “salary” indicando que no alcanza el mínimo imponible (el mensaje mostrado por consola deberá decir: “error: el mínimo imponible es de 150.000 y el salario ingresado es de: [salary]”, siendo [salary] el valor de tipo int pasado por parámetro).
*/

import (
	"fmt"
	"os"
)

func myCustomErrorTest(salary int) (int, error) {
	if salary < 150000 {
		return 400, fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salary)
	}
	return 200, nil
}

func main() {
	var salary int
	salary = 150000
	_, err := myCustomErrorTest(salary)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("Debe pagar impuesto")
}
