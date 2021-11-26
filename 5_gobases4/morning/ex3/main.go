package main

import "fmt"

type MyError struct {
	salario int
	msg	   string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Salario %d - %s", e.salario, e.msg)
}


func Ej1(salary int) (string, error) {

	if salary < 150000 {
		return "", fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salary)
	}
	return "debe pagar impuesto", nil
}

func main() {
	fmt.Println(Ej1(140000))
}