package main

import "fmt"

type myError struct {
	salario int
	msg     string
}

func (e *myError) Error() string {
	return "error: el salario ingresado no alcanza el mínimo imponible"
}

func verificarError(salario int) error {
	if salario < 150000 {
		return &myError{salario,
			""}
	}
	return nil
}

func main() {
	var salary = 100000
	err := verificarError(salary)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuesto")
	}

	salary2 := 160000
	err = verificarError(salary2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuesto")
	}

	// if salary < 150000 {
	// 	err := lanzarError()
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Debe pagar impuesto")
	// }
}
