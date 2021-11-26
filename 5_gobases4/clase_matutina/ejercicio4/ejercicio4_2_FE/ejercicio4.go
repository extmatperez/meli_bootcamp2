package main

import (
	"fmt"
)

func main() {
	salary, err := SalarioPorHoras(50, 50.55)
	if err != nil {
		fmt.Printf("Respuesta: %v\n", err)
	} else {
		fmt.Printf("Respuesta: %v\n", salary)
	}

	bonus, err := CalcularAguinaldo(7582.5, -6)
	if err != nil {
		fmt.Printf("Respuesta: %v\n", err)
	} else {
		fmt.Printf("Respuesta: %v\n", bonus)
	}

}

func SalarioPorHoras(horasTrabajadas int, valorHora float64) (float64, error) {
	salarioTotal := 0.0
	if horasTrabajadas >= 80 {
		salarioTotal = float64(horasTrabajadas) * valorHora
		if salarioTotal >= 150000 {
			salarioTotal -= salarioTotal * 0.1
		}
	} else {
		return 0.0, fmt.Errorf("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	}
	return salarioTotal, nil
}

func CalcularAguinaldo(mejorSalarioSem float64, mesesTrabajadosSem int) (float64, error) {
	if mejorSalarioSem < 0 || mesesTrabajadosSem < 0 {
		return 0.0, fmt.Errorf("No pueden haber valores negativos")
	} else {
		return (mejorSalarioSem / 12 * float64(mesesTrabajadosSem)), nil
	}
}
