package main

import (
	"errors"
	"fmt"
)

const (
	minimo   = "minimo"
	promedio = "promedio"
	maximo   = "maximo"
)

// ejercicio 1
func calculo_impuesto_porcentaje(salario float64, porcentaje float64) float64 {
	return salario * porcentaje / 100
}
func calcular_impuesto_salario(salario float64) float64 {
	switch {
	case salario > 150000.0:
		return calculo_impuesto_porcentaje(salario, 27)
	case salario > 50000.0:
		return calculo_impuesto_porcentaje(salario, 17)
	default:
		return 0
	}
}

// EJERCICIO 2
func verificar_numero(numero int) (float64, error) {
	if numero < 0 {
		return 0.0, errors.New("un numero es negativo")
	}
	return float64(numero), nil
}
func calcular_promedio_calificaciones(calificaciones ...int) (float64, error) {
	suma := 0.0
	for _, calif := range calificaciones {
		num, err := verificar_numero(calif)
		if err != nil {
			return 0, err
		} else {
			suma += num
		}
	}
	return suma / float64(len(calificaciones)), nil
}

//EJERCICIO 3
func calcular_sueldo_por_categoria(horas float64, sueldo_por_hora float64, porcentaje_adicional float64) float64 {
	subtotal := 0.0
	subtotal = horas * sueldo_por_hora
	return subtotal + subtotal*porcentaje_adicional/100
}
func calcular_salario_por_horas_trabajadas(minutos_trabajados int, categoria string) float64 {
	if categoria == "C" {
		return calcular_sueldo_por_categoria(float64(minutos_trabajados)/60.0, 1000.0, 0.0)
	} else if categoria == "B" {
		return calcular_sueldo_por_categoria(float64(minutos_trabajados)/60.0, 1500.0, 20.0)
	} else if categoria == "A" {
		return calcular_sueldo_por_categoria(float64(minutos_trabajados)/60.0, 3000.0, 50.0)
	} else {
		return 0
	}
}

//EJERCICIO 4

func operacion(tipo_operacion string) (func(...int) float64, error) {
	switch tipo_operacion {
	case minimo:
		return Min, nil
	case maximo:
		return Max, nil
	case promedio:
		return Prom, nil
	}
	return nil, errors.New("operacion no reconocida")
}
func Min(valores ...int) float64 {
	var minimo = valores[0]
	for _, num := range valores {
		if num < minimo {
			minimo = num
		}
	}
	return float64(minimo)
}
func Max(valores ...int) float64 {
	var maximo = valores[0]
	for _, num := range valores {
		if num > maximo {
			maximo = num
		}
	}
	return float64(maximo)
}
func Prom(valores ...int) float64 {
	var suma = 0
	for _, num := range valores {
		suma += num
	}
	return float64(suma) / float64(len(valores))
}

//EJERCICIO 5

func main() {
	fmt.Println("IMPUESTO: ", calcular_impuesto_salario(1000000))
	promediiio, _ := calcular_promedio_calificaciones(2, 3, 4, 5, 5)
	fmt.Println("promedio calificaciones: ", promediiio)
	fmt.Println("sueldo de categoria: ", calcular_salario_por_horas_trabajadas(480, "A"))
	minFunc, _ := operacion(minimo)
	// promFunc, _ := operacion(promedio)
	// maxFunc, _ := operacion(maximo)

	valorMinimo := minFunc(2, 3, 3, 4, 1, 2, 4, 5)
	fmt.Println("Minimo: ", valorMinimo)
}
