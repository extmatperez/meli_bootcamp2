package main

import "fmt"

func main() {

	months := map[int]string{1: "Enero", 2: "Febrero",
		3: "Marzo", 4: "Abril", 5: "Mayo", 6: "Junio", 7: "Julio", 8: "Agosto", 9: "Septiembre",
		10: "Octubre", 11: "Noviembre", 12: "Diciembre"}

	var numberOfMonth int
	fmt.Println("Introduce el numero del mes que quieres saber:")
	fmt.Scanf("%d\n", &numberOfMonth)

	fmt.Println("Tu mes es:", months[numberOfMonth])
}
