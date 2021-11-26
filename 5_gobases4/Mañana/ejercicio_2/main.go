package main

import (
	"errors"
	"fmt"
)

func main() {
	salary := 5

	if salary < 15000 {
		fmt.Println(errors.New("error: el salario ingresado no alcanza el mínimo imponible"))
	} else {
		fmt.Println("Correcto")

	}

}
