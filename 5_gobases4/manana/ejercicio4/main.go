package main

import (
	"errors"
	"fmt"
)

func calcSalario(numHoras, valxHora float64) (float64, error) {
	var salario float64 = 0.0
	var err error = nil

	salMax := 150000.0
	salario = numHoras * valxHora
	if salario >= salMax {
		salario -= salario * 0.1

	}
	if numHoras < 80 {
		err = errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")

	}
	return salario, err
}

func calcAguinaldo(mejSalario, mesesTrabajados float64) (float64, error) {
	var err3 error = nil
	aguinaldo := 0.0
	if mejSalario < 0 || mesesTrabajados < 0 {

		err := fmt.Errorf("\n || nivel 1 del error || no se aceptan valores negativos :  %.2f,  %.2f   ", mejSalario, mesesTrabajados)
		err2 := fmt.Errorf("\n || nivel 2 del error || %w    ", err)
		err3 = fmt.Errorf("\n || nivel 3 del error || %w    ", err2)
	} else {
		aguinaldo = mejSalario / (12 * mesesTrabajados)

	}
	return aguinaldo, err3
}

func main() {
	fmt.Println("---------ejercicio A")
	salario, err := calcSalario(79, 2500)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("salario :", salario)
	}

	fmt.Println("\n---------ejercicio B")
	aguinaldo, err2 := calcAguinaldo(150000, 12)
	if err2 != nil {
		//fmt.Println(errors.Unwrap(errors.Unwrap(err2)))
		//fmt.Println(errors.Unwrap(err2))
		fmt.Println(err2)

	} else {
		fmt.Println("aguinado :", aguinaldo)
	}

}
