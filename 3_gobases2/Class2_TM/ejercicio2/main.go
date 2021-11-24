package main

import "fmt"

/*
Un colegio de Buenos Aires necesita calcular el promedio (por alumno)
de sus calificaciones. Se solicita generar una función en la cual se
le pueda pasar N cantidad de enteros y devuelva el promedio y un error
en caso que uno de los números ingresados sea negativo
*/

func cant_ent() {
	//var n int
	var notas int
	//array := []int{}
	array := []int{1, 4, 6, 4, 7, 7, 5, 9}
	fmt.Printf("Ingrese la cantidad de notas a cargar: \n")
	for _, nota := range array {
		//fmt.Scanf("%d", &n)
		notas += nota
		//array = append(array, n)
	}

	fmt.Println(notas / len(array))
}

func average(values ...int) int {
	var resul int
	for _, value := range values {
		resul += value
	}
	return resul / (len(values))
}

func main() {
	cant_ent()
	avr := average(1, 4, 6, 4, 7, 7, 5, 9)
	fmt.Printf("\nEl promedio es: %d", avr)
}
