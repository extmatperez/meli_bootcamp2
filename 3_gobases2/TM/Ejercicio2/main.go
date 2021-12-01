/*Un colegio de Buenos Aires necesita calcular el promedio (por alumno) de sus calificaciones. Se solicita generar una función
en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio y un error en caso que uno de los números ingresados
sea negativo
*/
package main

import (
	"fmt"
)

func promedio(notas ...int) float64 {
	var promedio int
	var count float64 = 0

	for _, value := range notas {
		promedio += value
		count++
	}

	return float64(promedio) / count

}

func main() {
	var nota1, nota2, nota3, nota4 int
	nota1 = 5
	nota2 = 10
	nota3 = 7
	nota4 = 1
	fmt.Printf("El promedio es: %v\n", promedio(nota1, nota2, nota3, nota4))

}
