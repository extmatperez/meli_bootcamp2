package calculadora

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumar(t *testing.T) {
	num1 := 100
	num2 := 200
	resultadoEsperado := 300

	resultado := Sumar(num1, num2)

	assert.Equal(t, resultadoEsperado, resultado, fmt.Sprintf("La funcion sumar obtuvo %v y debia obtener %v", resultado, resultadoEsperado))
}

func TestRestar(t *testing.T) {
	num1 := 100
	num2 := 200
	resultadoEsperado := -100

	resultado := Restar(num1, num2)

	assert.Equal(t, resultadoEsperado, resultado, fmt.Sprintf("La funcion restar obtuvo %v y debia obtener %v", resultado, resultadoEsperado))
}
