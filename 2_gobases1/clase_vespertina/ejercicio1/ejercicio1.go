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
	fmt.Printf("La palabra '%v' tiene %v caracteres", palabraElegida, len(palabraElegida))
}

func imprimirCaracteres(palabraElegida string) {
	for _, v := range palabraElegida {
		fmt.Println(string(v))
	}
}
