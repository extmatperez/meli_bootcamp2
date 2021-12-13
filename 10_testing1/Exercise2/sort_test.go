package ordenar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortEquals(t *testing.T) {
	numbers := []int{5, 2, 4}
	expectResult := []int{2, 4, 5}

	results := Sort(numbers)

	assert.Equal(t, expectResult, results, "Be equals")
}

func TestSortNotEquals(t *testing.T) {
	numbers := []int{5, 2, 4, 1000, -2}
	expectResult := []int{5, 2, 4, 1000, -2}

	results := Sort(numbers)

	assert.NotEqual(t, expectResult, results, "Be not equals")
}
