package main

import (
	"errors"
	"fmt"
)

const (
	SalarioA, BonoA = 3000, 50
	SalarioB, BonoB = 1500, 20
	SalarioC, BonoC = 1000, 0
)

func main() {
	salario, err := calcularSalario(15020, "A")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El salario total es: %f", salario)
	}
}

func calcularSalario(minutosTrabajadosXMes int, categoria string) (float64, error) {
	var horas float64 = float64(minutosTrabajadosXMes / 60)

	switch categoria {
	case "C":
		return catC(horas), nil
		// return calcularPorCat(horas,SalarioC,BonoC),nil
	case "B":
		return catB(horas), nil

	case "A":
		return catA(horas), nil
	}

	return 0.0, errors.New("Solo se admiten categorias A, B, C y en mayusculas.")
}

// func calcularPorCat(horas float64 , salario , bono int) float64{
// 	salarioMensual := float64(salario) * horas

// 	if bono != 0 {
// 		return  salarioMensual + salarioMensual * bono / 100
// 	}else{
// 		return  salarioMensual
// 	}
// }

func catC(horas float64) float64 {
	return SalarioC * horas
}
func catB(horas float64) float64 {
	salarioMensual := SalarioB * horas
	return salarioMensual + salarioMensual*BonoB/100
}
func catA(horas float64) float64 {
	salarioMensual := SalarioA * horas
	return salarioMensual + salarioMensual*BonoA/100
}
