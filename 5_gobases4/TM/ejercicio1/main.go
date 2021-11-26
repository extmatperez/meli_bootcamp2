/*
En tu función “main”, define una variable llamada “salary” y
asignarle un valor de tipo “int”.
Crea un error personalizado con un struct que implemente “Error()” con
el mensaje “error: el salario ingresado no alcanza el mínimo imponible"
y lánzalo en caso de que “salary” sea menor a 150.000. Caso contrario,
imprime por consola el mensaje “Debe pagar impuesto”
*/

package main

import "fmt"

type errorSalary struct {
}

func (e *errorSalary) Error() string {
	return "error: el salario ingresado no alcanza el mínimo imponible"
}

func main() {
	salary := 100000
	err := CalcularImpuesto(salary)
	if err != nil {
		fmt.Println(err)
	}

}

func CalcularImpuesto(salary int) error {

	if salary < 150000 {
		return &errorSalary{}
	} else {
		fmt.Printf("Debe pagar impuesto\n")
		return nil
	}
}
