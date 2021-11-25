package main

import "fmt"

// Ejercicio 4 - Calcular estadísticas
// Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones de los alumnos de un curso, requiriendo calcular los valores mínimo, máximo y promedio de sus calificaciones.

// Se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio) y que devuelva otra función ( y un error en caso que el cálculo no esté definido) que se le puede pasar una cantidad N de enteros y devuelva el cálculo que se indicó en la función anterior
// Ejemplo:

// const (
//    minimo = "minimo"
//    promedio = "promedio"
//    maximo = "maximo"
// )

// ...

// minFunc, err := operacion(minimo)
// promFunc, err := operacion(promedio)
// maxFunc, err := operacion(maximo)

// ...

// valorMinimo := minFunc(2,3,3,4,1,2,4,5)
// valorPromedio := promFunc(2,3,3,4,1,2,4,5)
// valorMaximo := maxFunc(2,3,3,4,1,2,4,5)

func calculoMinimo(valores ...float64) float64 {
	min := valores[0]
	for i := 1; i < len(valores); i++ {
		if valores[i] < min {
			min = valores[i]
		}
	}
	return min
}

func calculoPromedio(valores ...float64) float64 {
	sum := 0.0
	lenght := float64(len(valores))

	for i := 0; i < len(valores); i++ {
		sum += valores[i]
	}

	return sum / lenght
}

func calculoMaximo(valores ...float64) float64 {
	max := valores[0]
	for i := 1; i < len(valores); i++ {
		if valores[i] > max {
			max = valores[i]
		}
	}
	return max

}

func estadistica(operacion string) func(valores ...float64) float64 {

	switch operacion {
	case "minimo":
		return calculoMinimo
	case "promedio":
		return calculoPromedio
	case "maximo":
		return calculoMaximo
	}
	return nil

}

func main() {

	// notas := []float64{5.5, 6.0, 10.5}

	tipoDeOp := estadistica("maximo")
	resultado := tipoDeOp(5.5, 10, 9.5, 8.5)

	fmt.Println(resultado)

}
