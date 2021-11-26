package main

import (
	"errors"
	"fmt"
)

type Empleado struct {
	Nombre            string
	Apellido          string
	DNI               int
	SalariosMensuales []float64
}

func main() {
	empleado := Empleado{"Nombre", "Apellido", 1231231, nil}
	agregarSalario(&empleado, 90, 150.0)
	agregarSalario(&empleado, 70, 150.0)
	agregarSalario(&empleado, 311, 200.0)
	agregarSalario(&empleado, 100, 155.0)
	agregarSalario(&empleado, 100, 190.0)
	agregarSalario(&empleado, 300, 190.0)
	agregarSalario(&empleado, 113, 170.0)
	agregarSalario(&empleado, 110, 190.0)
	agregarSalario(&empleado, 100, 150.0)
	agregarSalario(&empleado, 300, 190.0)
	//agregarSalario(&empleado, 300, -190.0)

	//empleado.SalariosMensuales = append(empleado.SalariosMensuales, -10000)

	aguinaldo, err := CalcularMedioAguinaldo(empleado)
	if err != nil {
		//fmt.Println(err)
		err2 := errors.Unwrap(err)
		fmt.Println(err2)
	} else {
		fmt.Printf("El aguinaldo es de: %.2f\n", aguinaldo)
	}
}

func CalcularSalarioMensual(horasTrabajadas int, valorHora float64) (float64, error) {
	if horasTrabajadas < 80 {
		err := errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
		return 0, fmt.Errorf("error al calcular el salario: %w", err)
	}
	salario := float64(horasTrabajadas) * valorHora
	if salario < 150000 {
		return salario, nil
	} else {
		return salario * 0.90, nil
	}
}

func CalcularMedioAguinaldo(empleado Empleado) (float64, error) {
	var salarios []float64
	var cantidad int
	if len(empleado.SalariosMensuales) >= 6 {
		cantidad = 6
	} else {
		cantidad = len(empleado.SalariosMensuales)
	}
	salarios = empleado.SalariosMensuales[len(empleado.SalariosMensuales)-cantidad:]

	fmt.Printf("Cantidad Salarios: %d\n", len(salarios))
	mejorSalario := 0.0
	for _, salario := range salarios {
		if salario < 0 {
			err := errors.New("Ningún salario no puede ser negativo")
			return 0, fmt.Errorf("error: %w", err)
		}
		if salario > mejorSalario {
			mejorSalario = salario
		}
	}
	return mejorSalario / 12 * float64(cantidad), nil
}

/*func CalcularMedioAguinaldo(empleado Empleado) (float64, error) {
	//var salarios []float64
	var cantidad int
	if len(empleado.SalariosMensuales) >= 6 {
		cantidad = 6
	} else {
		cantidad = len(empleado.SalariosMensuales)
	}
	//fmt.Printf("Cantidad Salarios: %d\n", len(salarios))
	mejorSalario := 0.0
	for _, salario := range empleado.SalariosMensuales[len(empleado.SalariosMensuales)-cantidad:] {
		if salario < 0 {
			err := errors.New("Ningún salario puede ser negativo")
			return 0, fmt.Errorf("error: %w", err)
		}
		if salario > mejorSalario {
			mejorSalario = salario
		}
	}
	return mejorSalario / 12 * float64(cantidad), nil
}*/

func agregarSalario(empleado *Empleado, horasTrabajadas int, salario float64) {
	sal, err := CalcularSalarioMensual(horasTrabajadas, salario)
	if err != nil {
		//fmt.Println(err)
		err2 := errors.Unwrap(err)
		fmt.Println(err2)
	} else {
		//fmt.Printf("Salario: %.2f\n", sal)
		empleado.SalariosMensuales = append(empleado.SalariosMensuales, sal)
	}
}
