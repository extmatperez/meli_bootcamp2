package main

import "fmt"

func main() {
	meses := [12]string{"Enero","Febrero","Marzo","Abril","Mayo","Junio","Julio","Agosto","Septiembre","Octubre","Noviembre","Diciembre"}
	mes := 2
	fmt.Println(meses[mes - 1])
}

/* Otra forma de resolverlo hubiera sido usar un switch con los 12 casos, uno por cada mes, esta forma quizas sea mas eficiente
en cuanto a uso de memoria y velocidad */