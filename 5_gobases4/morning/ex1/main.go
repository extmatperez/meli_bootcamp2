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
		return "", &MyError{
			salary,
			"error: el salario ingresado no alcanza el mínimo imponible",
		} 
	}
	return "debe pagar impuesto", nil
}

func main() {
	fmt.Println(Ej1(160000))
}