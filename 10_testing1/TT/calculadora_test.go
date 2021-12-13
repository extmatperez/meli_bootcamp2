package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumar(t *testing.T) {
	//Arrange
	num1 := 3
	num2 := 5

	resultadoEsperado := 8

	//Act
	resultado := Sumar(num1, num2)

	//Assert
	if resultado != resultadoEsperado {
		t.Errorf("La función Sumar obtuvo %v pero el resultado esperado era %v", resultado, resultadoEsperado)
	}
}

func TestSumar2(t *testing.T) {
	//Arrange
	num1 := 3
	num2 := 5

	resultadoEsperado := 8

	//Act
	resultado := Sumar(num1, num2)

	//Assert
	assert.Equal(t, resultadoEsperado, resultado)
}

func TestRestar(t *testing.T) {
	//Arrange
	num1 := 3
	num2 := 5

	resultadoEsperado := -2

	//Act
	resultado := Restar(num1, num2)

	//Assert
	if resultado != resultadoEsperado {
		t.Errorf("La función Restar obtuvo %v pero el resultado esperado era %v", resultado, resultadoEsperado)
	}
}
