package main

import (
	"testing"

	"github.com/stretchr/testify/assert" // Se importa testify
)

func TestRestar(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 10
	num2 := 3
	resultadoEsperado := 7

	// Se ejecuta el test
	resultado := Restar(num1, num2)

	// Se validan los resultados
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
}

func TestOrdenar(t *testing.T) {
	nums := []int{4, 1, 5, 2, 6, 9}

	ordenado := Ordenar(nums)
	esperado := []int{1, 2, 4, 5, 6, 9}
	assert.Equal(t, esperado, ordenado, "deben ser iguales")
}

func TestDividir(t *testing.T) {
	num := 6.0
	dem := 0.0
	esperado := 3.0

	division, err := Dividir(num, dem)
	if err != nil {
		if assert.Error(t, err) {
			assert.Equal(t, err.Error(), "no se puede dividir por 0")
		}
	} else {
		assert.Nil(t, err)
		assert.Equal(t, esperado, division, "deben ser iguales")
	}

}
