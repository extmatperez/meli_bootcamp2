package main

import (
	"errors"
	"fmt"
)

// type myError struct {
// 	status int
// 	msg    string
// }

// func (e *myError) Error() string {
// 	return "error: el salario ingresado no alcanza el mínimo imponible"
// }

// func test(n int) error {
// 	if n < 15000 {
// 		return &myError{

// 			msg: "error: el salario ingresado no alcanza el mínimo imponible",
// 		}
// 	}
// 	return nil
// }

func salarioMes(horas int, valor float64) (float64, error) {
	total := float64(horas) * valor
	if total >= 150000 {
		total -= total * 0.1
	}
	if (horas < 80) || (horas < 0) {
		return total, errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	}

	return total, nil
}

func calcAguinaldo(salario float64, meses int) (float64, error) {
	var aguinaldo float64
	if meses > 0 {
		aguinaldo = (salario / 12) * float64(meses)
		return aguinaldo, nil
	} else {
		err := fmt.Errorf("Error: no ha completado el tiempo de un mes para el aguinaldo")
		return aguinaldo, err
	}

}

func main() {
	salary := 50.00
	horas := 50
	meses := -1
	total, err := salarioMes(horas, salary)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(total)

	}
	agui, err2 := calcAguinaldo(salary, meses)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(agui)

	}

}
