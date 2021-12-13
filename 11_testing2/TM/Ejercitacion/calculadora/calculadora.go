package calculadora

import (
	"errors"
	"sort"
)

type Logger interface {
	Log(string) error
}

// Creado para el MOCK. Tambien es util para un FAKE para simular una entrada random de cierto dato.
type Config interface {
	SumaEnabled(cliente string) bool
}

// Funcion que recibe dos enteros, un objeto del tipo logger y retorna la suma resultante.
func SumarDummy(n1, n2 int, logger Logger) int {
	err := logger.Log("Ingreso a la función Sumar.")
	if err != nil {
		return -99999
	}
	return n1 + n2
}

// Funcion que recibe dos enteros y retorna la suma resultante con un cliente y un config como parametros para el MOCK.
func SumaRestricted(n1, n2 int, config Config, cliente string) int {
	if !config.SumaEnabled(cliente) {
		return -99999
	}
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
	return nums
}

// Funcion que ordena un slice de enteros de forma descendente.
func OrdenarDesc(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	return nums
}
