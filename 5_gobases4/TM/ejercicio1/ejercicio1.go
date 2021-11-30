package main

import "fmt"

type myError struct {
	msg string //
}

func (e *myError) Error() string {
	return fmt.Sprintf("Error: %s", e.msg)
}

func devolver_custom_error() error {
	var e myError
	e.msg = "El salario ingresado no alcanza el minimo imponible."
	return &e
}

func main() {
	var salary int
	fmt.Println("Ingresa el salario:")
	fmt.Scanf("%d", &salary)
	if salary < 150000 {
		err1 := devolver_custom_error()
		fmt.Println(err1)
	} else {
		fmt.Println("Debe pagar impuesto.")
	}
}
