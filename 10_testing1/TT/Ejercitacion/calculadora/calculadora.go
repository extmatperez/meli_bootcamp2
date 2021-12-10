package calculadora

import (
	"errors"
	"fmt"
	"sort"
)

// Funcion que recibe dos enteros y retorna la suma resultante.
func Sumar(n1, n2 int) int {
	return n1 + n2
}

// Funcion que recibe dos enteros y retorna la resta o diferencia resultante.
func Restar(n1, n2 int) int {
	return n1 - n2
}

// Funcion que recibe dos enteros y retorna el producto resultante.
func Multiplicar(n1, n2 int) int {
	return n1 * n2
}

// Funcion que recibe dos enteros y retorna el resultado de la división exacta resultante.
func Dividir(n1, n2 int) (int, error) {
	if n2 == 0 {
		return 0, errors.New("El denominador no puede ser igual a cero.")
	}
	return n1 / n2, nil
}

// Funcion que ordena un slice de enteros de forma ascendente.
func OrdenarAsc(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for _, v := range nums {
		fmt.Println(v)
	}
	return nums
}

// Funcion que ordena un slice de enteros de forma descendente.
func OrdenarDesc(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	for _, v := range nums {
		fmt.Println(v)
	}
	return nums
}
