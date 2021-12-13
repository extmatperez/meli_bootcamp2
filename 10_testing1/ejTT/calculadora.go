package calculadora

import "sort"

func Sumar(n1, n2 int) int {
	return n1 + n2
}

func Restar(n1, n2 int) int {
	return n1 - n2
}

func Ordenar(valores []int) []int {
	sort.Ints(valores)
	return valores

}
