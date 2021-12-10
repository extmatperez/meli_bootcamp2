package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiv(t *testing.T) {
	num1 := 3.0
	num2 := 3.0
	esperado := 1

	resl, err := Dividir(float64(num1), float64(num2))
	assert.Nil(t, err)
	assert.Equal(t, float64(esperado), float64(resl), "Valor esperado: %v \nValor obtenido: %v", esperado, resl)
}
func TestDiv2(t *testing.T) {
	num1 := 3.0
	num2 := 0

	_, err := Dividir(float64(num1), float64(num2))
	assert.NotNil(t, err)
}
