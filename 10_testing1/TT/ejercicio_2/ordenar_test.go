package ordenar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenar(t *testing.T) {
	arr := []int{6, 8, 3, 0, 7, 11, 2}
	esperado := []int{0, 2, 3, 6, 7, 8, 11}

	resl := Ordenar(arr)
	assert.Equal(t, esperado, resl, "Valor esperado: %v \nValor obtenido: %v", esperado, resl)
}
