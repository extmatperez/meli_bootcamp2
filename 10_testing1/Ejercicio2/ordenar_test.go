package ordenar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenar(t *testing.T) {
	a := []int{5, 3, 4, 7, 8, 9}
	slice := Ordenar(a)
	assert.IsIncreasing(t, slice, "slice ordenado")
}
