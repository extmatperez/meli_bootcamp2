package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResta(t *testing.T) {
	num1 := 3.0
	num2 := 3.0
	esperado := 0

	resl := Restar(float64(num1), float64(num2))
	assert.Equal(t, float64(esperado), float64(resl), "Valor esperado: %v \nValor obtenido: %v", esperado, resl)
}
