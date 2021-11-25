package main

import "fmt"

// Ejercicio 2 - Promedio
// Un colegio de Buenos Aires necesita calcular el promedio (por alumno) de sus calificaciones. Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio y un error en caso que uno de los números ingresados sea negativo

func promedio(scores ...float64) float64 {
	total_score := 0.0
	score_quantity := float64(len(scores))
	var average_score float64

	for _, score:= range scores{
		total_score = total_score + score
	}

	average_score = total_score/score_quantity

	return average_score
}


func main () {

	fmt.Println(promedio(10.0, 9.0, 8))

}
