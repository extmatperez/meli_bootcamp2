package main

import (
	"fmt"
)

func checkSalary(salary int) {

	if salary < 150000 {
		err := fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de:  %d", salary)
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuestos")
	}
}
func main() {

	salary := 10000
	checkSalary(salary)

}
