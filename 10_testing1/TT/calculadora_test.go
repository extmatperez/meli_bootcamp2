package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumar(t *testing.T) {
	num1 := 3
	num2 := 5
	resultadoEsperado := 8
	resultado := Sumar(num1, num2)
	//
	//if resultado != resultadoEsperado {
	//	t.Errorf("Funcion suma() arrojo el resultado= %v , pero el esperado es %v", resultado, resultadoEsperado)
	//}
	//con testify
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
}

func TestDividir(t *testing.T) {
	num1 := 3
	num2 := 3
	resultado := Dividir(num1, num2)
	assert.NotNil(t, resultado)
}
