package gotesting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestar(t *testing.T) {
	a := 8
	b := 2
	resEsperado := 6

	resActual := Restar(a, b)

	assert.Equal(t, resEsperado, resActual)
}

func TestOrdenar(t *testing.T) {

	var slic = []int{1, 2, 5, 3, 4}
	resActual := Ordenar(slic)

	assert.IsIncreasing(t, resActual)
}

func TestDividir(t *testing.T) {
	num := 6
	den := 0
	resEsperado := 2

	resObtenido, err := Dividir(num, den)

	assert.Nil(t, err, "Error encontrado")
	assert.Equal(t, resEsperado, resObtenido, "No son iguales")
}
