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
	fmt.Println()
	salary, err := myCustomErrorTest(160.000)
	if err != nil { // si err no es nulos es por que existe un error
		fmt.Println(err)
	} else {
		fmt.Println(salary)
	}
	fmt.Println()
}
