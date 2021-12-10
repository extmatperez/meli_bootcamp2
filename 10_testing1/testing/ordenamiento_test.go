package testing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenar(t *testing.T) {
	//Inicializar
	arreglo := []int{5, 7, 8, 3, 1}
	resultadoEsperado := []int{1, 3, 5, 7, 8}

	//Llamado a función
	SortByInsertion(arreglo)

	//Assert
	assert.Equal(t, resultadoEsperado, arreglo, "Debían ser iguales")
}

func TestOrdenarFallido(t *testing.T) {
	//Inicializar
	arreglo := []int{5, 7, 8, 3, 1}
	resultadoEsperado := []int{1, 3, 5, 8}

	//Llamado a función
	SortByInsertion(arreglo)

	//Assert
	assert.NotEqual(t, resultadoEsperado, arreglo, "No debían ser iguales")
}
