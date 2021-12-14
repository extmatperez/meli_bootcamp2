package fibonacci

func fibo(num int) int {
	if num < 2 {
		return num
	}

	a1 := 0
	a2 := 1
	for i := 1; i < num; i++ {
		a2, a1 = a2+a1, a2
	}

	return a2
}
