package go_testing

import (
	"fmt"
	"testing"
)

func TestIntsArrayOrder(t *testing.T) {
	ints := 			[]int{9,23,16,39,19,47,10,1}
	expectedResult := 	[]int{1,9,10,16,19,23,39,47}

	//result := IntArrayOrder(ints)
	result := orderSlices(ints)
	fmt.Println(ints)
	fmt.Println(expectedResult)
	if !Equal(ints, expectedResult) {
		t.Errorf("Funcion IntsArrayOrder(ints []int]) []int => Arrojo el resultado %v, pero el resultado esperado es %v", result, expectedResult)
	}
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}