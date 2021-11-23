package main

import "fmt"

func main() {
	var numMes int = 1
	aQueMesCorresponde(numMes)

	numMes = 2
	aQueMesCorresponde(numMes)

	numMes = 3
	aQueMesCorresponde(numMes)

	numMes = 8
	aQueMesCorresponde(numMes)
}

func aQueMesCorresponde(numeroDelMes int) {
	meses := map[int]string{}
	meses[1] = "Enero"
	meses[2] = "Febrero"
	meses[3] = "Marzo"
	meses[4] = "Abril"
	meses[5] = "Mayo"
	meses[6] = "Junio"
	meses[7] = "Julio"
	meses[8] = "Agosto"
	meses[9] = "Septiembre"
	meses[10] = "Octubre"
	meses[11] = "Noviembre"
	meses[12] = "Diciembre"

	fmt.Println(meses[numeroDelMes])
}

// Otra solucion que se me ocurre es un switch que dependiendo el case , imprima el mes. Pero creo que seria muuuy largo
