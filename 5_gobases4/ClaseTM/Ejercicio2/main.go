package main

import (
	"errors"
	"fmt"
)

type myError struct {
	msg string
}

// func (e *myError) Error() string {
// 	return e.msg
// }

func crearError(salario int) (string, error) {
	if salario < 150000 {
		return "", errors.New("error: el salario ingresado no alcanza el mínimo imponible")
	} else {
		return "Debe pagar impuesto", nil
	}
}

func main() {
	var salario int = 200220

	ok, err := crearError(salario)

	if err == nil {
		fmt.Println(ok)
	} else {
		fmt.Println(err)
	}
}
