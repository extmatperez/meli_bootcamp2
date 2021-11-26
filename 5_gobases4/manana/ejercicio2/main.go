package main

import (
	"errors"
	"fmt"
)

func checkSalary(salary int) {

	if salary < 150000 {
		err := errors.New("error: el salario ingresado no alcanza el mínimo imponible")
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuestos")
	}
}
func main() {

	salary := 150000
	checkSalary(salary)

}
