package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestar(t *testing.T) {
	//Arrange
	num1 := 5
	num2 := 6
	resultadoEsperado := -1

	//Act
	resultado := Restar(num1, num2)

	//Assert
	assert.Equal(t, resultado, resultadoEsperado)
}
