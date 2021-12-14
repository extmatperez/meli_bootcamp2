package fibonacci

func fibonacciRecursiva(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 1
	default:
		return fibonacciRecursiva(n-1) + fibonacciRecursiva(n-2)
	}
}

func fibonacciIterativa(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	n_1 := 1
	n_2 := 1
	resultado := 2
	for i := 3; i <= n; i++ {
		resultado = n_1 + n_2
		n_2 = n_1
		n_1 = resultado
	}
	return resultado
}
