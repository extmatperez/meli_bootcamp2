package main

import "fmt"

func main()  {
	
	fmt.Println("Arrancamos con el ejercicio 6!!!")

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	//PARTE 1
	fmt.Println("La edad de Benjamin es", employees["Benjamin"])

	//PARTE 2
	var contador uint
	contador = 0

	for _, element := range employees {
		if element > 21 {
			contador++
		}
	}

	fmt.Println("Los empleados mayor a 21 son", contador)

	//PARTE 3
	fmt.Println("Antes los empleados eran en total", len(employees))
	employees["Federico"] = 25
	fmt.Println("Ahora con Fede son en total", len(employees))

	//PARTE 4
	delete(employees, "Pedro")
	fmt.Println("Pedro se fue a una empresa mas grande, ahora son", len(employees))
}