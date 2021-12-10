package ordenamiento

import (
	"sort"
)

type Nums struct {
	num1 int
	num2 int
	num3 int
	num4 int
}

func (n Nums) Order(resultado ...int) []int {

	resultado = append(resultado, n.num1, n.num2, n.num3, n.num4)
	sort.Ints(resultado)
	return resultado
}
