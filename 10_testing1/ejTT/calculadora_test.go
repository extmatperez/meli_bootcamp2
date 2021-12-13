package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculadora(t *testing.T) {

	n1, n2 := 3, 4
	resultadoEsperadoSuma := 7
	resultadoEsperadoResta := -1
	var resultadoEsperadoOrdenarArray []int
	resultadoEsperadoOrdenarArray = append(resultadoEsperadoOrdenarArray, 0, 1, 2, 3, 5, 6, 7, 8)
	arrayEnteros := []int{2, 6, 8, 3, 5, 7, 1, 0}
	// arrayEnteros = append(arrayEnteros, 2, 6, 8, 3, 5, 7, 1, 0)

	resultadoSuma := Sumar(n1, n2)
	resultadoResta := Restar(n1, n2)
	arrayOrdenado := Ordenar(arrayEnteros)

	assert.Equal(t, resultadoEsperadoSuma, resultadoSuma, "deberian ser iguales")
	assert.Equal(t, resultadoEsperadoResta, resultadoResta, "deberian ser iguales")
	assert.NotEqual(t, resultadoEsperadoResta, resultadoResta+1, "deberian ser iguales")
	assert.Equal(t, resultadoEsperadoOrdenarArray, arrayOrdenado, "deberian estar ordenados")
	assert.IsIncreasing(t, arrayOrdenado, "debe estar ordenado ascendentemente")
	// if resultado != resultadoEsperado {
	// 	t.Errorf("function TestSumar FAILED, expect %v, got %v", resultadoEsperado, resultado)
	// }
}
