/*
Repite el proceso anterior, pero ahora implementando “fmt.Errorf()”,
para que el mensaje de error reciba por parámetro el valor de “salary”
indicando que no alcanza el mínimo imponible (el mensaje mostrado por consola
deberá decir: “error: el mínimo imponible es de 150.000 y el salario ingresado es
de: [salary]”, siendo [salary] el valor de tipo int pasado por parámetro).
*/
package main

import (
	"fmt"
)

func testSalary3(value int) error {
	if value < 150000 {
		return fmt.Errorf("Error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", value)
	} else {
		return nil
	}
}

func main() {
	var salary1 int = 160000
	err := testSalary3(salary1)
	if err != nil {
		fmt.Println("Error ocurrido: ", err)
	} else {
		fmt.Printf("Salario ingresado: %d \nResultado: Debe pagar impuesto\n", salary1)
	}

}
