package testing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestar(t *testing.T) {
	//Inicializar
	num1 := 5
	num2 := 3
	resultadoEsperado := 2

	//Llamado a función
	resultado := Restar(num1, num2)

	//Assert
	assert.Equal(t, resultadoEsperado, resultado, "Debían ser iguales")
}

func TestRestarFallido(t *testing.T) {
	//Inicializar
	num1 := 8
	num2 := 3
	resultadoEsperado := 2

	//Llamado a función
	resultado := Restar(num1, num2)

	//Assert
	assert.NotEqual(t, resultadoEsperado, resultado, "No debían ser iguales")
}

func TestDividir(t *testing.T) {
	//Inicializar
	num1 := 6
	num2 := 3
	resultadoEsperado := 2

	//Llamado a función
	resultado, err := Dividir(num1, num2)

	//Assert
	assert.Equal(t, resultadoEsperado, resultado, "Debían ser iguales")
	assert.Nil(t, err, "Debía ser nil")
}

func TestDividirCero(t *testing.T) {
	//Inicializar
	num1 := 5
	num2 := 0
	resultadoEsperado := 0

	//Llamado a función
	resultado, err := Dividir(num1, num2)

	//Assert
	assert.Equal(t, resultadoEsperado, resultado, "Debían ser iguales")
	assert.NotNil(t, err, "No debía ser nil")
}
