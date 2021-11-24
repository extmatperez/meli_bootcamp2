package main

import "fmt"

func main() {
	/*Se me ocurren dos maneras de realizarlo. La primera es utilizando 12 if para ver cual es el mes, y la otra es utilizando
	  un mapa que asocie cada número del mes con su nombre */
	mapMonth := map[int]string{1: "Enero", 2: "Febrero", 3: "Marzo", 4: "Abril",
		5: "Mayo", 6: "Junio", 7: "Julio", 8: "Agosto",
		9: "Septiembre", 10: "Octubre", 11: "Noviembre", 12: "Diciembre"}

	monthNum := 6
	monthName := mapMonth[monthNum]

	if monthNum >= 1 && monthNum <= 12 {
		fmt.Printf("El mes correspondiente al número %v es \"%v\"\n", monthNum, monthName)
	} else {
		fmt.Printf("El mes %v no existe.\n", monthNum)
	}

}
