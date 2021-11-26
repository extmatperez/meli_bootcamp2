package main

import "fmt"

type myCustomError struct {
	status int
	msg    string
}

func (e *myCustomError) Error() string {

	return "Error creado 2" + e.msg

}

func myCustomErrorSalary(status int) error {

	if status == 1 {
		return &myCustomError{
			status: status,
			msg:    "error: el salario ingresado no alcanza el mínimo imponible"}

	}
	return nil
}
func checkSalary(salary int) {

	if salary < 150000 {
		err := myCustomErrorSalary(1)
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuestos")
	}
}
func main() {

	salary := 1000
	checkSalary(salary)

}
