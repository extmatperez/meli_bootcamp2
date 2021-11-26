/*
Haz lo mismo que en el ejercicio anterior pero reformulando el código para que, en reemplazo
de “Error()”, se implemente “errors.New()”.
*/

package main

import (
	"errors"
	"fmt"
)

func main() {

	salary := 100000

	if salary < 150000 {
		fmt.Println(errors.New("error: el salario ingresado no alcanza el mínimo imponible"))
	} else {
		fmt.Println(errors.New("debe pagar impuesto"))
	}
}
