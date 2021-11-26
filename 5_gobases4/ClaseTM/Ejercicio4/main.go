package main

import (
	"errors"
	"fmt"
)

func calcularSalario(horas int, valorHora float64) (float64, error) {
	if horas > 80 {
		salarioCalculado := float64(horas) * valorHora

		if salarioCalculado >= 150000 {
			salarioCalculado -= (salarioCalculado * 10 / 100)
		}
		return float64(salarioCalculado), nil
	} else {
		return 0, errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	}
}

func calcularMedioAguinaldo(mejorSalario float64, mesesTrabajados int) (float64, error) {
	if mejorSalario < 0 || mesesTrabajados < 0 {
		return 0, fmt.Errorf("No puede ingresar numeros negativos\nMejor Salario %f\nMeses Trabajados: %d", mejorSalario, mesesTrabajados)
	} else {
		return mejorSalario / 12 * float64(mesesTrabajados), nil
	}
}

func main() {
	horas := 15000
	preciHora := 100.59
	salario, err := calcularSalario(horas, preciHora)

	if err == nil {
		fmt.Printf("Horas trabajadas %d \nPrecio de hora %f \nSalario total: %f \n", horas, preciHora, salario)
	} else {
		fmt.Println(err)
	}

	aguinaldo, err := calcularMedioAguinaldo(150000, -4)
	if err == nil {
		fmt.Printf("El aguinaldo es: %.2f \n", aguinaldo)
	} else {
		fmt.Println(err)
	}

}
