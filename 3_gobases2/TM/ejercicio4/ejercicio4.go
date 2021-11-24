package main

import "fmt"
import "errors"
	/*
		Los profesores de una universidad de Colombia necesitan calcular algunas 
		estadísticas de calificaciones de los alumnos de un curso, requiriendo calcular 
		los valores mínimo, máximo y promedio de sus calificaciones.

Se solicita generar una función que indique qué tipo de cálculo se quiere realizar 
(mínimo, máximo o promedio) y que devuelva otra función ( y un error en caso que el 
	cálculo no esté definido) que se le puede pasar una cantidad N de enteros y devuelva 
	el cálculo que se indicó en la función anterior
Ejemplo:
const (
	minimo = "minimo"
	promedio = "promedio"
	maximo = "maximo"
)

...

minFunc, err := operacion(minimo)
promFunc, err := operacion(promedio)
maxFunc, err := operacion(maximo)

...

valorMinimo := minFunc(2,3,3,4,1,2,4,5)
valorPromedio := promFunc(2,3,3,4,1,2,4,5)
valorMaximo := maxFunc(2,3,3,4,1,2,4,5)
	*/

const (
	minimo = "minimo"
	promedio = "promedio"
	maximo = "maximo"
)

func minFunc(n ...float64) float64{
	aux:=n[0]
	for _, i := range n {
		if i < aux {
			aux = i
		}
	}
	return aux
}
func promFunc(n ...float64) float64{
	aux:=n[0]
	for _, i := range n {
		aux+=i
	}
	return aux/len(n)
}
func maxFunc(n ...float64) float64{
	aux:=n[0]
	for _, i := range n {
		if i > aux {
			aux = i
		}
	}
	return aux
}

func operacion(operacion string) func(n ...float64) {
switch operacion{
	case minimo:
		return minFunc
	case promedio:
		return promFunc1
	case maximo:
		return maxFunc
	default:
		return nil
}
}

minFunc, err := operacion(minimo)
promFunc, err := operacion(promedio)
maxFunc, err := operacion(maximo)


func main() {
	valorMinimo := minFunc(2,3,3,4,1,2,4,5)
	valorPromedio := promFunc(2,3,3,4,1,2,4,5)
	valorMaximo := maxFunc(2,3,3,4,1,2,4,5)

	fmt.Printf("Minimo: %v \n",valorMinimo)
	fmt.Printf("Promedio: %v \n",valorPromedio)
	fmt.Printf("Maximo: %v \n",valorMaximo)
}
