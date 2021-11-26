/*
Haz lo mismo que en el ejercicio anterior pero reformulando el código para que,
en reemplazo de “Error()”,  se implemente “errors.New()”.
*/
package main

import (
	"errors"
	"fmt"
)

func testSalary2(value int) {
	if value < 150000 {
		fmt.Printf("Salario ingresado: %d \n", value)
		fmt.Println(errors.New("El salario ingresado no alcanza el mínimo imponible"))
		return
	} else {
		fmt.Printf("Salario ingresado: %d \n", value)
		fmt.Println("Debe pagar impuesto")
		return
	}
}

func main() {
	var salary1 int = 100000
	testSalary2(salary1)

	var salary2 int = 160000
	testSalary2(salary2)

}
