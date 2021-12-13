package ordenamiento

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Ordenar(t *testing.T) {

	resultado := OrdenarNumeros(3, 4, 2, 7, 5)
	esperado := []int{2, 3, 4, 5, 7}

	assert.Equal(t, esperado, resultado, "Deben ser iguales")
}
