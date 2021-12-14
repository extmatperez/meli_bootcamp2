package ordenamiento

import "testing"

func testOrder(t testing.T) {
	ns := []int{1, 4, 6, 7, 2, 5}
	expectedResult := []int{1, 2, 4, 5, 6, 7}
	result := AscendingOrder(ns)
	err := false

	for i, _ := range result {
		if result[i] != expectedResult[i] {
			err = true
			break
		}
	}
	if err {
		t.Errorf("The order function should answer %v and it answered with %v", result, expectedResult)
	}
}
