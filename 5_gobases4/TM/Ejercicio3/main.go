package main

import (
	"fmt"
)

func funcError(salary int) error {
	if salary < 150000 {
		return fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salary)
	}
	return nil
}

func salarioMensual(horas int, valor_hora float64) (float64, error) {
	if horas < 80 {
		return 0, fmt.Errorf("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	}

	salario := float64(horas) * valor_hora

	return salario, nil
}

func main() {
	horas := 150
	valor := 150.0
	salario, err := salarioMensual(horas, valor)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(salario)
	}
}
