package main

import "fmt"

type MyError struct {
	salary int
	msg    string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Salario %d - %s", e.salary, e.msg)
}

func Ej1(salary int) (string, error) {
	if salary < 150000 {
		return "", &MyError{salary, "error: el salario ingresado no alcanza el mínimo imponible"}
	}
	return "Debe pagar impuestos", nil
}

func main() {
	fmt.Println(Ej1(1000))
}
