package main

import "fmt"

func main() {

	//PRACTICA CON PUNTEROS
	fmt.Println("\n--------PUNTEROS---------\n")

	text := "hola mundo"

	var ptext *string

	ptext = &text

	fmt.Println(*ptext, &ptext, ptext)
	fmt.Println(&text, text)

	fmt.Println("\n")

	//si cambio el valor de texto, el apuntador ptext cambia tambien
	text = "chao mundo"

	fmt.Println(*ptext, &ptext, ptext)
	fmt.Println(&text, text)

	//ARRAYS
	fmt.Println("\n--------ARRAYS---------\n")

	var array [4]int

	array[0] = 1
	array[1] = 2
	array[2] = 3
	array[3] = 4

	fmt.Println(array)
	fmt.Println(array[2])
	fmt.Println("\n")

	fmt.Println("\n--------SLICES---------\n")

	var mySlice []string

	mySlice = append(mySlice, "hola")
	mySlice = append(mySlice, "mundo")

	fmt.Println(mySlice)

	fmt.Println("\n")

	fmt.Println(mySlice[0])
	fmt.Println(mySlice[1])

	fmt.Println("\n")

	//delete(mySlice, 0)

	fmt.Println(mySlice)

	fmt.Println("\n")

	//-----------CREAR UN SLICE CON MAKE-----------

	a := make([]string, 4)

	fmt.Println(len(a))

	primes := []int{2, 3, 5, 7, 11}

	fmt.Println(primes[1:4])

	fmt.Println(len(primes), cap(primes))

	primes = append(primes, 17)

	fmt.Println(len(primes), cap(primes))

	fmt.Println("\n------------MAPS------------------\n")

	//myMap := map[string]int{}

	//myMap2 := make(map[string]string)

	var students = map[string]int{"Ariel": 3, "Juan": 4, "Pedro": 5}

	fmt.Println(students["Ariel"])
	fmt.Println(students)

	delete(students, "Ariel")

	fmt.Println(students)

	students["Ariel"] = 10 //ya existe el key, se reemplaza el valor
	students["Pablo"] = 4  //no existe el key, se agrega

	for key, value := range students {
		fmt.Println("La clave es: ", key, ". El valor es: ", value)
	}
}
