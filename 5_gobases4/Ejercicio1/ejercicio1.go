package main

import "fmt"

type myCustomError struct {
	msn string
}

func (e *myCustomError) Error() string {

	return fmt.Sprintf(e.msn)
}

func myCustomErrorTest(sal float64) (string, error) {
	if sal <= 150.000 {
		return "", &myCustomError{msn: "error: el salario ingresado no alcanza el minimo imponible"}
	}
	return "debe pagar impuestos", nil
}
func main() {

	salary, err := myCustomErrorTest(150.000)
	if err != nil { // si err no es nulos
		fmt.Println(err)
	} else {
		fmt.Println(salary)
	}
}
