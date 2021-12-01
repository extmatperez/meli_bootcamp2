/* Una empresa de inteligencia artificial necesita tener una funcionalidad para crear una estructura que represente una matriz
de datos.
Para ello requieren una estructura Matrix que tenga los métodos:
Set:  Recibe una serie de valores de punto flotante e inicializa los valores en la estructura Matrix
Print: Imprime por pantalla la matriz de una formas más visible (Con los saltos de línea entre filas)
La estructura Matrix debe contener los valores de la matriz, la dimensión del alto, la dimensión del ancho, si es cuadrática y
cuál es el valor máximo.
*/

package main

import "fmt"

type Matrix struct {
	Alto       int
	Ancho      int
	Cuadratica bool
	Maximo     int
	Valores    []float64
}

func (v *Matrix) set(valores ...float64) {
	v.Valores = valores
}
func (v *Matrix) print() {
	alto := 0
	ancho := 0
	for _, value := range v.Valores {
		if ancho == v.Ancho {
			fmt.Println()
			ancho = 0
			alto++
		}
		if alto == v.Alto {
			break
		}
		fmt.Printf("%.1f\t", value)
		ancho++
	}
}

func main() {
	matriz1 := Matrix{
		Alto:       4,
		Ancho:      4,
		Cuadratica: true,
		Maximo:     25,
	}
	matriz1.set(5, 4, 7, 15, 14, 25, 24, 19, 11, 1, 23, 12, 9, 13, 6, 7)
	fmt.Println("\nMATRIZ DE DATOS:")
	matriz1.print()
	fmt.Println()
}
