package main

import "fmt"

type errorPersonalizado struct {
}

func (e errorPersonalizado) Error() string {
	return "error: el salario ingresado no alcanza el mínimo imponible"
}

func pagaImpuesto(salary int) (string, error) {
	miError := errorPersonalizado{}

	if salary < 150000 {
		return "", miError
	} else {
		return "Debe pagar impuesto", nil
	}

}

func main() {

	salary := 150000

	respuesta, err := pagaImpuesto(salary)

	if err != nil {
		fmt.Printf("Ocurrio un error %v\n", err.Error())
	} else {
		fmt.Println(respuesta)
	}

}