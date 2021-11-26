package main

import (
	"errors"
	"fmt"
	"os"
)

type myCustomError struct {
	status int
	msg    string
}

func getError() error {
	return errors.New("Esto es un error")
}

func (e *myCustomError) Error() string {
	return "Error"
}

func myCustomErrorTest(status int) (int, error) {
	if status >= 300 {
		return 400, &myCustomError{
			status: status,
			msg:    "algo salió mal",
		}
	}
	return 200, nil
}

func main() {

	err := getError()
	fmt.Println(err)

	status, err := myCustomErrorTest(250) //llamamos a nuestra func
	if err != nil {                       //hacemos una validación del valor de err
		fmt.Println(err) //si err no es nil, imprimimos el error y...
		os.Exit(1)       //utilizamos este método para salir del programa
	}
	fmt.Printf("Status %d, Funciona!", status)
}
