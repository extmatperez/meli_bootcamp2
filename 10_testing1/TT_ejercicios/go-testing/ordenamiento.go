package calculadora

import (
	"sort"
)

func Ordenar(desordenado []int) []int {
	sort.Ints(desordenado)
	return desordenado
}
