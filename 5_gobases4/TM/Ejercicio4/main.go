package main

import (
	"fmt"
)

func calcularSalario(horas, precioHoras int) (salarioCalculado float64, err error) {

	if horas < 80 || horas < 0 {
		err := fmt.Errorf("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
		fmt.Println(err)
	} else {
		salarioCalculado = float64(horas * precioHoras)
		err = fmt.Errorf("")
	}

	if salarioCalculado >= 150000 {
		salarioCalculado = salarioCalculado * 0.9
	}

	return salarioCalculado, err
}


func calcularAguinaldo(salarios[]float64) float64 {


}

func main() {
	salarios := [6]float64
	var horas int
	var precioHoras int

	fmt.Println("Ingrese la cantidad de horas trabajadas")
	fmt.Scanln(&horas)
	fmt.Println("Ingrese el precio por hora")
	fmt.Scanln(&precioHoras)
	salario, err := calcularSalario(horas, precioHoras)
	salario.

	fmt.Printf("el salario del trabajador es", salario)
	fmt.Println(err)

}
