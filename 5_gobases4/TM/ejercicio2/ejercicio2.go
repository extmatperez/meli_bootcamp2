package main

import (
	"errors"
	"fmt"
	"os"
)

/*
Haz lo mismo que en el ejercicio anterior pero reformulando el código para que, en reemplazo de “Error()”,  se implemente “errors.New()”.
*/

func myCustomErrorTest(salary int) (int, error) {
	if salary < 150000 {
		return 400, errors.New("error: el salario ingresado no alcanza el mínimo imponible")
	}
	return 200, nil
}

func main() {
	var salary int
	salary = 140000
	_, err := myCustomErrorTest(salary)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("Debe pagar impuesto")
}
