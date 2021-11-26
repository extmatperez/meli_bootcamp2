package main

import (
	"fmt"
	"os"
)

/*
En tu función “main”, define una variable llamada “salary” y asignarle un valor de tipo “int”.
Crea un error personalizado con un struct que implemente “Error()” con el mensaje “error: el salario ingresado no alcanza el mínimo imponible" y lánzalo en caso de que “salary” sea menor a 150.000. Caso contrario, imprime por consola el mensaje “Debe pagar impuesto”.
*/

type myCustomError struct {
	status int
	msg    string
}

func (e *myCustomError) Error() string {
	return fmt.Sprintf("status:%d, %s\n", e.status, e.msg)
}

func myCustomErrorTest(salary int) (int, error) {
	if salary < 150000 {
		return 400, &myCustomError{
			status: 400,
			msg:    "error: el salario ingresado no alcanza el mínimo imponible",
		}
	}
	return 200, nil
}

func main() {
	var salary int
	salary = 150000
	_, err := myCustomErrorTest(salary)
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}
	fmt.Println("Debe pagar impuesto")
}
