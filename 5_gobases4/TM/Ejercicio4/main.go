package main

import (
	"errors"
	"fmt"
)

func salarioMensual(horas int, valor_hora float64) (float64, error) {
	if horas < 80 {
		return 0, fmt.Errorf("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	}

	salario := float64(horas) * valor_hora

	if salario >= 150000 {
		salario -= (salario * 0.1)
	}

	return salario, nil
}

func aguinaldo(mejor_salario, meses_trabajados int) (float64, error) {
	if mejor_salario < 0 || meses_trabajados < 0 {
		return 0, errors.New("error: ha ingresado un valor negativo")
	}
	aguinaldo := (float64(mejor_salario) / 12) * float64(meses_trabajados)
	return aguinaldo, nil
}

func main() {
	horas := 1
	valor := 1500.0
	salario, err := salarioMensual(horas, valor)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(salario)
	}

	aguinaldo, err := aguinaldo(2000, -6)
	if err != nil {
		fmt.Println(err)
	} else {

		fmt.Println(aguinaldo)
	}

}
