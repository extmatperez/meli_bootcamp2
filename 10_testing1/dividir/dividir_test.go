package dividir

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividir(t *testing.T) {
	resul, err := Dividir(10, 5)

	assert.Nil(t, err)
	assert.Equal(t, 2, resul, "Deben ser iguales")
}
func TestDividirFailed(t *testing.T) {
	_, err := Dividir(10, 0)

	assert.NotNil(t, err)
	//assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
}
