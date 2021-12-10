package ordenamiento

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenamiento(t *testing.T) {
	slice := []int{1, 8, 3, 10, 58, 2}
	expectedSlice := []int{1, 2, 3, 8, 10, 58}

	result := OrderingSlice(slice)

	assert.Equal(t, expectedSlice, result, "Slice is not sorted properly")
}