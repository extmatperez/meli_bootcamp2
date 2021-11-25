package main

import "fmt"

/*
Realizar una aplicación que contenga una variable con el número del mes.
Según el número, imprimir el mes que corresponda en texto.
¿Se te ocurre si se puede resolver de más de una manera? ¿Cuál elegirías y por qué?
Ej: 7, Julio
*/
func main() {
	var meses = map[int]string{1: "Enero", 2: "Febrero", 3: "Marzo", 4: "Abril", 5: "Mayo", 6: "Junio", 7: "Julio", 8: "Agosto", 9: "Septiembre", 10: "Octubre", 11: "Noviembre", 12: "Diciembre"}

	var mes int = 1
	for mes > 0 && mes < 13 {
		fmt.Printf("Ingrese el número de mes: ")
		_, err := fmt.Scanf("%d\r\n", &mes)
		if err == nil {
			fmt.Printf("El mes es: %s\n", meses[mes])
		} else {
			mes = -1
		}
	}

}
