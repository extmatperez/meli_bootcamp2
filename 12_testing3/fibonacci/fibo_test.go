package fibonacci

import (
	"testing"
)

func TestFiboZero(t *testing.T) {
	//Init
	tests := []struct {
		arg  int
		want int
	}{{0, 0}, {1, 1}, {2, 1}, {3, 2}, {7, 13}, {9, 34}}

	for i, d := range tests {
		got := fibo(d.arg)
		if got != d.want {
			t.Errorf("Test[%d]: fibonacci(%d) returned %d, wanted %d", i, d.arg, got, d.want)
		}
	}

}
