package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	arrayToOrder := []int{2, 6, 8, 10, 4, 9}
	arraySpected := []int{2, 4, 6, 8, 9, 10}
	resultado := Sort(arrayToOrder)
	assert.Equal(t, resultado, arraySpected, "El array no esta ordenado")

}
