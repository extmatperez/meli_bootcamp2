package main

import (
	"errors"
	"fmt"
)

func main() {
	/*
		Haz lo mismo que en el ejercicio anterior pero reformulando el código para que,
	en reemplazo de “Error()”,  se implemente “errors.New()”.
	 */

	var (
		salary = 160000
	)
	err := checker(salary)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Su salario si alcanza el minimo imponible")
	}


}
func checker(salary int) error {
	if salary < 150000 {
		return errors.New("Su salario no alcanza el minimo impoible")
	}
	return nil
}
