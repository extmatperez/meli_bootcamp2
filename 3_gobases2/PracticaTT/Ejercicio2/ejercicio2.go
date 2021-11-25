package main

import "fmt"

/*
Una empresa de inteligencia artificial necesita tener una funcionalidad para crear una estructura que represente
 una matriz de datos.
Para ello requieren una estructura Matrix que tenga los métodos:
Set:  Recibe una serie de valores de punto flotante e inicializa los valores en la estructura Matrix
Print: Imprime por pantalla la matriz de una formas más visible (Con los saltos de línea entre filas)
La estructura Matrix debe contener los valores de la matriz, la dimensión del alto, la dimensión del ancho,
si es cuadrática y cuál es el valor máximo.
*/

type Matrix struct {
	valores    []float64
	alto       int
	ancho      int
	cuadratica bool
	max        float64
}

func (m *Matrix) Set(valores ...float64) {
	m.valores = valores

	//Verifico si es cuadratica
	if m.ancho == m.alto {
		m.cuadratica = true
	} else {
		m.cuadratica = false
	}

	//Seteo el maximo
	m.max = m.valores[0]
	for _, valor := range m.valores {
		if valor > m.max {
			m.max = valor
		}
	}
}

func (m Matrix) Print() {
	alturaAct := 0
	anchoAct := 0
	for _, valor := range m.valores {
		fmt.Printf("%v\t", valor)
		if m.comprobarLimites(&alturaAct, &anchoAct) == -1 {
			break
		}
	}
}

func (m Matrix) comprobarLimites(alturaAct *int, anchoAct *int) int {
	//Verifico el ancho para hacer el salto de linea en el print
	//Tambien verifico el alto. Si me paso del alto dejo de imprimir
	*anchoAct++
	if *anchoAct == m.ancho {
		fmt.Printf("\n")
		*alturaAct++
		*anchoAct = 0
		if *alturaAct == m.alto {
			return -1
		}
	}
	return 0
}

func main() {

	matriz1 := Matrix{alto: 3,
		ancho: 3,
	}
	matriz1.Set(100, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3)
	fmt.Printf("%v\n", matriz1)
	matriz1.Print()

}
