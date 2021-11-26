package main

import "fmt"

func main() {
	var alto, ancho, cuadratica, maximo int

	fmt.Println("Ejericio 2")
	fmt.Println("Ingrese alto")
	fmt.Scanf("%v", &alto)
	fmt.Println("Ingrese ancho")
	fmt.Scanf("%v", &ancho)
	fmt.Println("Ingrese cuadratica")
	fmt.Scanf("%v", &cuadratica)
	fmt.Println("Ingrese maximo")
	fmt.Scanf("%v", &maximo)

	matrix := Matrix{}
	matrix.Set(float64(alto), float64(ancho), float64(cuadratica), float64(maximo))
	matrix.Print()

}

type Matrix struct {
	Alto, Ancho, Cuadratica, Maximo float64
}

func (m *Matrix) Set(valores ...float64) {
	m.Alto = valores[0]
	m.Ancho = valores[1]
	m.Cuadratica = valores[2]
	m.Maximo = valores[3]
}

func (m Matrix) Print() {
	fmt.Printf("Alto: %v \nAncho: %v \nCuadratica: %v \nMaxValor: %v \n", m.Alto, m.Ancho, m.Cuadratica, m.Maximo)
}
