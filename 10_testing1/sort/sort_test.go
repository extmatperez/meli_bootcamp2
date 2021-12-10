package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortAscending(t *testing.T) {
	nums := []int{10, 15, 20, 1, 3, 2, 5}
	expectedResult := []int{1, 2, 3, 5, 10, 15, 20}
	actualResult := SortAscending(nums)

	assert.Equal(t, expectedResult, actualResult, "must be equal")
}
