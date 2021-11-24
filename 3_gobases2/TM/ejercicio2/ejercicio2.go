package main

import "fmt"

func promedio(notas ...int) {
	/*
		Un colegio de Buenos Aires necesita calcular el promedio (por alumno)
		de sus calificaciones. Se solicita generar una función en la cual se le
		pueda pasar N cantidad de enteros y devuelva el promedio y un error en
		caso que uno de los números ingresados sea negativo
	*/
	total := 0.0
	for _, i := range notas {
		if i < 0 {
			fmt.Printf("ERROR, no ingrese numeros negativos\n")
		} else {
			total += float64(i)
		}
	}
	fmt.Printf("El promedio es: %v \n", total/float64(len(notas)))
}

func main() {
	promedio(1, 5, 8, 7, 6, 7, 9, 8, 7, 5)
}
