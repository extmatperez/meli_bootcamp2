package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumar(t *testing.T) {

	//Arrenge
	num1 := 3
	num2 := 10
	resultadoEsperado := 13

	//Act
	resultado := Sumar(num1, num2)

	//Assert
	if resultadoEsperado != resultado {
		t.Errorf("El resultado de la función sumar fue: %v, pero el resultado esperado era: %v", resultado, resultadoEsperado)
	}
}

func TestRestar(t *testing.T) {

	//Arrenge
	num1 := 3
	num2 := 10
	resultadoEsperado := -7

	//Act
	resultado := Restar(num1, num2)

	//Assert
	if resultadoEsperado != resultado {
		t.Errorf("El resultado de la función restar fue: %v, pero el resultado esperado era: %v", resultado, resultadoEsperado)
	}
}

// Con testfy
func TestSumarTestify(t *testing.T) {

	//Arrenge
	num1 := 3
	num2 := 10
	resultadoEsperado := 13

	//Act
	resultado := Sumar(num1, num2)

	//Assert
	assert.Equal(t, resultadoEsperado, resultado)
}

func TestDividirTestify(t *testing.T) {

	//Arrenge
	num1 := 3
	num2 := 1

	//Act
	resultado := Dividir(num1, num2)

	//Assert
	assert.NotNil(t, resultado)
}
