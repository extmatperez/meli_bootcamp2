package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenar(t *testing.T) {
	//Arrange
	lista := []int{5, 4, 7, 1, -2, 23, -8, 10}
	resultadoEsperado := []int{-8, -2, 1, 4, 5, 7, 10, 23}

	//Act
	Ordenar(lista)

	//Assert
	assert.Equal(t, lista, resultadoEsperado, "deben ser iguales")
}

func TestOrdenarSort(t *testing.T) {
	//Arrange
	lista := []int{5, 4, 7, 1, -2, 23, -8, 10}
	resultadoEsperado := []int{-8, -2, 1, 4, 5, 7, 10, 23}

	//Act
	OrdenarSort(lista)

	//Assert
	assert.Equal(t, lista, resultadoEsperado, "deben ser iguales")
}
