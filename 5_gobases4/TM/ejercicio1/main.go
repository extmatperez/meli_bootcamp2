/*
En tu función “main”, define una variable llamada “salary” y asignarle un valor de tipo “int”.

Crea un error personalizado con un struct que implemente “Error()” con el mensaje “error: el
salario ingresado no alcanza el mínimo imponible" y lánzalo en caso de que “salary” sea menor
a 150.000. Caso contrario, imprime por consola el mensaje “Debe pagar impuesto”.
*/

package main

import "fmt"

type myError struct {
	msg string
}

func (e *myError) Error() string {

	return e.msg
}

func myErrorTest(salary int) error {

	if salary < 150000 {
		return &myError{"error: el salario ingresado no alcanza el mínimo imponible"}
	}

	return &myError{"Debe pagar impuesto"}
}

type error interface {
	Error() string
}

func main() {

	salary := 500000

	err := myErrorTest(salary)

	fmt.Println(err)
}
