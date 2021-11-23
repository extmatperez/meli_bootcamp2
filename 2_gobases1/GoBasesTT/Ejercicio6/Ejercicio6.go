package main

import "fmt"

func main() {
	queEdadTiene("Benjamin")
}

func queEdadTiene(nombre string) {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Println("La edad de ", nombre, "es ", employees[nombre])

	var mayoresA21 int = 0
	for _, elem := range employees {
		if elem > 21 {
			mayoresA21++
		}
	}

	fmt.Println("Los mayores de 21 son : ", mayoresA21)

	employees["Federico"] = 25
	delete(employees, "Pedro")
	fmt.Println("Despues de agregar a Federico y quitar a Pedro, quedaron los siguientes: ", employees)

}
