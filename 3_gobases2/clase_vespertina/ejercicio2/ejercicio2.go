package main

import "fmt"

func main() {
	notas := leerNotas()
	calcularPromedio(notas)
}

func leerNotas() []float64 {
	var notas []float64
	fmt.Println("Para salir escriba -1 \n")
	for {
		nota := capturarNotas()
		if nota != -1 {
			notas = append(notas, nota)
		} else {
			break
		}
		fmt.Printf("\n")
	}
	return notas
}

func capturarNotas() float64 {
	var nota float64
	fmt.Println("Ingrese la nota:")
	fmt.Scanf("%f", &nota)
	return nota
}

func calcularPromedio(notas []float64) {
	cantidad := len(notas)
	totalNotas := 0.0
	for _, val := range notas {
		totalNotas += val
	}
	total := totalNotas / float64(cantidad)
	fmt.Printf("El promedio es de %v \n", total)
}
