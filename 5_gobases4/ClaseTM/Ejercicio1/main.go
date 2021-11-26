package main

import "fmt"

type myError struct {
	msg string
}

func (e *myError) Error() string {
	return e.msg
}

func main() {
	var salario int = 200200

	if salario < 150000 {
		err := &myError{msg: "error: el salario ingresado no alcanza el mínimo imponible"}
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}
