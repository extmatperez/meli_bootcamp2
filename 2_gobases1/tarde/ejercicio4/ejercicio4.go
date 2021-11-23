package main

import "fmt"

func mesCorrespondiente(numeroMes int) {
	var meses = []string{"enero", "febrero", "marzo", "abril", "mayo",
		"junio", "julio", "agosto", "septiembre", "octubre", "noviembre", "diciembre"}
	if numeroMes > len(meses) || numeroMes <= 0 {
		fmt.Println("El numero excede la cantidad de meses del año")
	} else {
		fmt.Printf("El numero corresponde al mes %v", meses[numeroMes-1])
	}
}

func main() {
	mesCorrespondiente(1)
}
