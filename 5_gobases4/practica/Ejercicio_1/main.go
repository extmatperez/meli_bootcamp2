package main

import "fmt"

type MyError struct {
	msg 	string
	status	int
}

func main() {
	/*
		En tu función “main”, define una variable llamada “salary” y asignarle un valor de tipo “int”.
		Crea un error personalizado con un struct que implemente “Error()” con el mensaje “error: el
		salario ingresado no alcanza el mínimo imponible" y lánzalo en caso de que “salary” sea menor
		a 150.000. Caso contrario, imprime por consola el mensaje “Debe pagar impuesto”.
	 */
	var (
		salary = 20000
	)
	err := checker(salary)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Su salario si alcanza para ser imponible.")
	}
}
func checker(salary int) error {
	if salary < 150000 {

		return &MyError{"Algo salio mal", 400}
	}
	return nil
}
func (e MyError) Error() string {
	return "error: El salario ingresado no alcanza el minimo imponible"
}