package main

import "fmt"

func main() {

	mes := 1

	meses := [12]string{
		"enero",
		"febrero",
		"marzo",
		"abril",
		"mayo",
		"junio",
		"julio",
		"agosto",
		"septiembre",
		"octubre",
		"noviembre",
		"diciembre",
	}

	fmt.Printf("El mes es %s\n", meses[mes-1])
}
