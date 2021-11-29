package main

import (
	"errors"
	"fmt"
)

func calcularSalario(horas, valorHora int) (int, error) {

	if horas < 80 {
		return 0, errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	} else {
		return horas * valorHora, nil
	}

}

func pagaImpuesto(salary int, e error) (int, error) {

	if salary <= 0 {
		return 0, fmt.Errorf("El salario debe ser mayor a 0: %w", e)
	}

	if salary < 150000 {
		return salary, e
	} else {
		return salary - ((salary * 10) / 100), e
	}

}

func calcularAguinaldo(mejorSalario int, meses int) (int, error) {
	if mejorSalario < 0 {
		return 0, errors.New("El salario no puede ser negativo")
	}
	if meses < 0 {
		return 0, fmt.Errorf("Los meses no pueden ser negativos")
	}

	if meses == 0 {
		return 0, fmt.Errorf("Los meses tienen que ser mayor a 0")
	}

	return mejorSalario / 12 * meses, nil
}

func main() {

	horas := 80
	valorhora := 2000

	salary, err := calcularSalario(horas, valorhora)
	if err != nil {
		//Error de horas
		fmt.Println(err.Error())
	}

	salaryMenosImpuesto, err2 := pagaImpuesto(salary, err)

	if err2 != nil {
		//error de salario
		fmt.Println(err2.Error())
		//error por el cual hubo error en el salario
		fmt.Println(errors.Unwrap(err2))
	} else {
		fmt.Printf("\n El salario es de %d \n", salaryMenosImpuesto)
	}

	aguinaldo, err := calcularAguinaldo(150000, 1)

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Printf("\nEl aguinaldo es de: %d\n", aguinaldo)
	}

}
