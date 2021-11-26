/*
1- En tu función “main”, define una variable llamada “salary” y asignarle un valor de tipo “int”.
2- Crea un error personalizado con un struct que implemente “Error()” con el mensaje
“error: el salario ingresado no alcanza el mínimo imponible" y lánzalo en caso de que
“salary” sea menor a 150.000. Caso contrario, imprime por consola el mensaje “Debe pagar impuesto”.
*/
package main

import (
	"fmt"
)

type myCustomError struct {
	status int
	msg    string
}

func (e *myCustomError) Error() string {
	return fmt.Sprintf("Error: %v", e.msg)
}

func testSalary(value int) (int, error) {
	if value < 150000 {
		return 500, &myCustomError{
			status: 500,
			msg:    "El salario ingresado no alcanza el mínimo imponible",
		}
	}
	return 200, nil
}

func main() {
	var salary1 int = 100000

	_, err1 := testSalary(salary1)
	if err1 != nil {
		fmt.Printf("Salario ingresado: %d \nResultado: %s \n", salary1, err1)
	} else {
		fmt.Printf("Salario ingresado: %d \nResultado: Debe pagar impuesto\n", salary1)
	}

	var salary2 int = 160000
	_, err2 := testSalary(salary2)
	if err2 != nil {
		fmt.Printf("Salario ingresado: %d \nResultado: %s\n", salary2, err2)
	} else {
		fmt.Printf("Salario ingresado: %d \nResultado: Debe pagar impuesto\n", salary2)
	}

}
