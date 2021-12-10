package calculadora

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividir(t *testing.T) {
	//Arrange
	num := 5
	den := 0
	resultadoEsperado := 0
	errorEsperado := errors.New("el denominador no puede ser 0")

	//Act
	resultado, err := Dividir(num, den)

	//Assert
	if err != nil {
		assert.Error(t, errorEsperado, err)
	} else {
		assert.Equal(t, resultadoEsperado, resultado)
	}
}
