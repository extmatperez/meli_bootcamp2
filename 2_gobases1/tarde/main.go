package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	//Ejercicio 1
	reader := bufio.NewReader(os.Stdin)
	var palabra string
	fmt.Println("Ingrese la palabra: ")
	palabra, _ = reader.ReadString('\n')
	var letras []string = strings.Split(palabra, "")
	fmt.Printf("La palabra tiene %d letras\n", len(letras))
	fmt.Print("Palabra deletreada: ")
	for _, letra := range letras {
		fmt.Printf("%s ", letra)
	}
	fmt.Println()

	//Ejercicio 2
	var precio, descuento float32
	fmt.Println("Ingrese el precio: ")
	fmt.Scanf("%f", &precio)
	fmt.Println("Ingrese el descuento en porcentaje: ")
	fmt.Scanf("%f", &descuento)

	precioFinal := (precio * descuento) / 100
	fmt.Printf("El precio final es de %v pesos", precioFinal)
}
