package main

import (
	"errors"
	"fmt"
)

func obtenerSalario(horas int, valorHora int) (float64, error) {
	if horas < 80 {
		err := fmt.Errorf("hubo un error: %w", errors.New("el trabajador no puede haber trabajado menos de 80 hs mensuales"))
		return 0, errors.Unwrap(err)
	}

	salario := float64(horas * valorHora)
	if salario >= 150000 {
		salario *= 0.9
	}

	return salario, nil
}

func obtenerAguinaldo(mejorSalario int, mesesTrabajados int) (int, error) {
	if mesesTrabajados < 0 {
		err := fmt.Errorf("hubo un error: %w", errors.New("los meses trabajados no pueden ser negativos"))
		return 0, errors.Unwrap(err)
	}
	if mejorSalario < 0 {
		err := fmt.Errorf("hubo un error: %w", errors.New("el mejor salario no puede ser negativo"))
		return 0, errors.Unwrap(err)
	}
	aguinaldo := mejorSalario / 12 * mesesTrabajados
	return aguinaldo, nil
}

func main() {
	salario, err := obtenerSalario(50, 2000)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El salario es de ", salario)
	}

	salario, err = obtenerSalario(150, 2000)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El salario es de ", salario)
	}

	aguinaldo, err := obtenerAguinaldo(160000, -2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El aguinaldo es ", aguinaldo)
	}

	aguinaldo, err = obtenerAguinaldo(-160000, 10)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El aguinaldo es ", aguinaldo)
	}

	aguinaldo, err = obtenerAguinaldo(160000, 10)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El aguinaldo es ", aguinaldo)
	}
}
