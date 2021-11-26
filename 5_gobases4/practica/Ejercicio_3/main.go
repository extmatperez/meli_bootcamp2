package main

import "fmt"

func main() {
	/*
		Repite el proceso anterior, pero ahora implementando “fmt.Errorf()”, para que el mensaje de error
		reciba por parámetro el valor de “salary” indicando que no alcanza el mínimo imponible (el mensaje
		mostrado por consola deberá decir: “error: el mínimo imponible es de 150.000 y el salario ingresado
		es de: [salary]”, siendo [salary] el valor de tipo int pasado por parámetro).
	 */
	var (
		salary =140000
	)

	err := checker(salary)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Su salario si alcanza el minimo imponible")
	}


}
func checker(salary int) error {
	if salary < 150000 {
		return fmt.Errorf("error: El minimo imponible es de 150.000 y el salario ingresado es: %v", salary)
	}
	return nil
}