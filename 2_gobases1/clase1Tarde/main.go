package main

import "fmt"

func firstExercise() {
	fmt.Println("Ejercicio 1")
	word := "paralelepipedo"

	fmt.Printf("La palabra tiene %v letras", len(word))
	fmt.Println("Sus letras son")
	for i := 0; i < len(word); i++ {
		fmt.Printf("%c ", word[i])
	}
	fmt.Println("")

}

func secondExercise() {
	fmt.Println("Ejercicio 2")
	price := 250.0
	discount := 30

	fmt.Printf("El precio con descuento es $%v\n", price*(1.0-float64(discount)/100))
}

func thirdExercise() {
	fmt.Println("Ejercicio 3")

	clientAge := 20
	clientIsEmployeed := true
	clientEmploymentAge := 3
	clientSalary := 9999999.99

	if clientAge < 22 || !clientIsEmployeed || clientEmploymentAge <= 1 {
		fmt.Println("No cumple los requisitos para otorgarle el credito.")
	} else if clientSalary < 100000.00 {
		fmt.Println("Se cobraran intereses por el credito")
	} else {
		fmt.Println("No se cobraran intereses por el credito")
	}
}

func fourthExercise() {
	fmt.Println("Ejercicio 4")

	month := 11
	switch month {
	case 1:
		fmt.Println("Enero")
	case 2:
		fmt.Println("Febrero")
	case 3:
		fmt.Println("Marzo")
	case 4:
		fmt.Println("Abril")
	case 5:
		fmt.Println("Mayo")
	case 6:
		fmt.Println("Junio")
	case 7:
		fmt.Println("Julio")
	case 8:
		fmt.Println("Agosto")
	case 9:
		fmt.Println("Septiembre")
	case 10:
		fmt.Println("Octubre")
	case 11:
		fmt.Println("Noviembre")
	case 12:
		fmt.Println("Diciembre")
	}

	// otras soluciones:
	// utilizar un map con key el indice del mes y value el nombre del mes {1: "Enero", 2: "Febrero"...}
	// utilizar un array de len 13, poniendo el [0] en un valor que simbolice error y utilizar los indicies 1-12 para buscar los meses.
	// eligiria el map.
}

func fifthExercise() {
	fmt.Println("Ejercicio 5")

	students := []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernán", "Leandro", "Eduardo", "Duvraschka"}

	newStudentsList := append(students, "Gabriela")
	fmt.Println(newStudentsList)
}

func sixthExercise() {
	fmt.Println("Ejercicio 6")

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Printf("La edad de Benjamin es: %v\n", employees["Benjamin"])
	fmt.Println("Los empleados mayor a 21 son:")

	for name, age := range employees {
		if age > 21 {
			fmt.Println(name)
		}
	}
	fmt.Println("Agregnado a Federico")
	employees["Federico"] = 25

	fmt.Println("Eliminando a Pedro")
	delete(employees, "Pedro")
}

func main() {
	firstExercise()
	secondExercise()
	thirdExercise()
	fourthExercise()
	fifthExercise()
	sixthExercise()
}
