package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortEqual(t *testing.T) {
	// (input/output)
	numbers := []int{5, 1, -1, 4, 0, 0, 2, 10, 10, 10000, -1000}
	expectedResult := []int{-1000, -1, 0, 0, 1, 2, 4, 5, 10, 10, 10000}

	// Execute test
	result := Sort(numbers)

	//check
	assert.Equal(t, expectedResult, result, "they should be equal")
}

func TestSortNotEqual(t *testing.T) {
	// (input/output)
	numbers := []int{5, 1, 10000, -1000}
	expectedResult := []int{5, 1, 10000, -1000}

	// Execute test
	result := Sort(numbers)

	//check
	assert.NotEqual(t, expectedResult, result, "they should be not equal")
}
