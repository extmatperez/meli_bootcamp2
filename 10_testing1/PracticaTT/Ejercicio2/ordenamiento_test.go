package testing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenamiento(t *testing.T) {
	//Arrange
	var slice []int
	slice = append(slice, 7, 5, 3, 6, 9)

	var sliceResultadoEsperado []int
	sliceResultadoEsperado = append(sliceResultadoEsperado, 3, 5, 6, 7, 9)

	//Act

	resultado := OrdenamientoSliceAscendente(slice)
	//Assert

	assert.Equal(t, resultado, sliceResultadoEsperado)
}
