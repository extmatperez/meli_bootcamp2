package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ingrese el numero del mes:")
	var numeroMes uint8 = 0
	fmt.Scanf("%d", &numeroMes)
	obtenerMesPorNumero(numeroMes)
}

func obtenerMesPorNumero(numeroMes uint8) {
	meses := map[uint8]string{
		1:  "Enero",
		2:  "Febrero",
		3:  "Marzo",
		4:  "Abril",
		5:  "Mayo",
		6:  "Junio",
		7:  "Julio",
		8:  "Agosto",
		9:  "Septiembre",
		10: "Octubre",
		11: "Noviembre",
		12: "Diciembre",
	}
	if _, ok := meses[numeroMes]; !ok {
		fmt.Println("El mes no existe")
	} else {
		fmt.Println(meses[numeroMes])
	}
}
