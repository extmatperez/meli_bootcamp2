package main

import (
	"errors"
	"fmt"
)

func main() {
	var salary int
	fmt.Println("Ingresa el salario:")
	fmt.Scanf("%d", &salary)
	if salary < 150000 {
		//err1 := devolver_custom_error()
		//fmt.Println(err1.msg)
		fmt.Println(errors.New("El salario ingresado no alcanza el minimo imponible."))
	} else {
		fmt.Println("Debe pagar impuesto.")
	}
}
