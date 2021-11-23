package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ingrese la palabra deseada:")
	var palabraElegida string
	fmt.Scanf("%s", &palabraElegida)
	fmt.Println("\n")
	contarCaracteres(palabraElegida)
	fmt.Println("\n")
	imprimirCaracteres(palabraElegida)
	fmt.Println("\n")
}

func contarCaracteres(palabraElegida string) {
	contador := 0
	for indice := range palabraElegida {
		contador = indice + 1
	}
	fmt.Println("La palabra", palabraElegida, "tiene", contador, "caracteres")
}

func imprimirCaracteres(palabraElegida string) {
	for _, v := range palabraElegida {
		fmt.Println(string(v))
	}
}
