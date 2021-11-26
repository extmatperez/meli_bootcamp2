/*
? Ejercicio 1 - Impuestos de salario #1
En tu función “main”, define una variable llamada “salary” y asignarle un valor de tipo “int”.
Crea un error personalizado con un struct que implemente “Error()” con el mensaje
“error: el salario ingresado no alcanza el mínimo imponible" y lánzalo en caso de que
“salary” sea menor a 150.000. Caso contrario, imprime por consola el mensaje “Debe pagar impuesto”.
*/

package main

import "fmt"

type myCustomError struct {
	msg string
}

func (e *myCustomError) Error() string {
	return fmt.Sprintf("%v", e.msg)
}

func myCustomErrorTest(value int) (int, error) {
	if value <= 150000 {
		return value, &myCustomError{msg: "el salario ingresado no alcanza el mínimo imponible"}
	}
	return value, nil
}

func main() {
	_, err := myCustomErrorTest(100000)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuestos")
	}

	_, err = myCustomErrorTest(200000)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuestos")
	}
}
