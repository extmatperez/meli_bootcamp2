package main

import (
	"fmt"
	"reflect"
)

func main() {
	var mes int

	fmt.Println("Ejercicio 4")
	fmt.Println("Ingrese numero del mes a evaluar")
	fmt.Scanf("%v", &mes)
	fmt.Println(mesesDelAnio(mes))

}

func mesesDelAnio(mes int) string {

	meses := map[int]string{1: "enero", 2: "febrero", 3: "marzo", 4: "abril", 5: "mayo", 6: "junio", 7: "julio", 8: "agosto", 9: "septiembre", 10: "octubre", 11: "noviembre", 12: "diciembre"}

	if reflect.ValueOf(mes).Kind() == reflect.Int {
		return meses[mes]
	} else {
		return " tenes que ingresar el numero del mes en cuestion "
	}

}
