package main

import (
	"fmt"
	"os"
)

func main() {
	num := 64

	fmt.Printf("%d %c %o %b %x %X %v %p", num, num, num, num, num, num, &num, &num)
	fmt.Printf("%-10d - %10d \n", num, num)
	fmt.Printf("%-10.2f - %10.10f\n", 129.11948, 125.444)

	//Sprintf

	salida := fmt.Sprintf("%b", 12456)
	fmt.Printf(salida)

	//setenv

	var err error
	err = os.Setenv("NAME", "jose")
	fmt.Print(err)
	err = os.Setenv("NAME", "juan")

	variable := os.Getenv("NAME")
	fmt.Println(variable)
	fmt.Println(err)

	variable2, ok := os.LookupEnv("NAdME")
	if ok {
		fmt.Printf("\nSe encontro la variable de entorno: %s", variable2)
	} else {
		fmt.Println("\nNo se encontro la variable de entorno")
	}

	data, err := os.ReadFile("/Users/joserios/Desktop/bootcamp/meli_bootcamp2/4_gobases3/a.txt")
	if err == nil {
		file := string(data)
		fmt.Println(file)
	} else {
		fmt.Println("El archivo no existe")
	}

	// WriteFile

}
