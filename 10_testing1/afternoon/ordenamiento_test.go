package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert" // Se importa testify
)

func TestOrdenar(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	lista := []int{5, 4, 3, 32, 11, 72, 9, 43, 86, 1, -3, -62}
	resultadoEsperado := []int{-62, -3, 1, 3, 4, 5, 9, 11, 32, 43, 72, 86}
	// Se ejecuta el test
	Ordenar(lista)
	// Se validan los resultados
	assert.Equal(t, resultadoEsperado, lista, "deben ser iguales")
}
