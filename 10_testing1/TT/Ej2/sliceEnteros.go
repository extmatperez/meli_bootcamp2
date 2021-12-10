//? Ejercicio 2 - Test Unitario Método Ordenar
// Diseñar un método que reciba un slice de enteros y los ordene de forma ascendente,
// posteriormente diseñar un test unitario que valide el funcionamiento del mismo.
// Dentro de la carpeta go-testing crear un archivo ordenamiento.go con la función a probar.
// Dentro de la carpeta go-testing crear un archivo ordenamiento_test.go con el test diseñado.

package slice_enteros

import (
	"sort"
)

func Slice(arg []int) []int {

	var contenedor []int
	sort.Slice(arg, func(i, j int) bool {
		return arg[i] < arg[j]
	})

	contenedor = append(contenedor, arg...)

	return contenedor
}
