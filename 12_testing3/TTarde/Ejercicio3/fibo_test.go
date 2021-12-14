package fibo

import "testing"

func TestFibo(t *testing.T) {
	tests := []struct {
		arg int
		want int
	}{{2, 1}}

	for i, d := range tests {
		got := fibo(d.arg)
		if got != d.want {
			t.Errorf("Test[%d]: fibo(%d) returned %d, want %d", i, d.arg, got, d.want)
		}
	}

	tests2 := []struct {
		arg int
		want int
	}{{0, 0}}

	for i, d := range tests2{
		got := fibo(d.arg)
		if got != d.want {
			t.Errorf("Test[%d]: fibo(%d) returned %d, want %d", i, d.arg, got, d.want)
		}
	}

}