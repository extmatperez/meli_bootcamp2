/*
Un colegio de Buenos Aires necesita calcular el promedio (por alumno) de sus calificaciones.
Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio y un error en caso que uno de los números ingresados sea negativo
*/
package main

import "fmt"

func average_calculations(ratings ...int) (float64, string) {
	var message string = "OK"
	var average float64 = 0.0
	var count_prom float64 = 0.0
	for count, rating := range ratings {
		if rating < 0 {
			message = "ERROR"
		}
		average += float64(rating)
		count_prom = float64(count)
	}

	return average / (count_prom + 1), message
}

func main() {
	average, message := average_calculations(100, 100, -60, 75)
	if message == "OK" {
		fmt.Printf("El promedio resultante es: %f", average)
	} else {
		fmt.Printf("%s: Se ingreso un valor negativo", message)
	}

}
