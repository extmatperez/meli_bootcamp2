package fibo

func fibo(x int) int{
	f := make([]int, x+1, x+2)
    if x < 2 {
        f = f[0:2]
    }
    f[0] = 0
    f[1] = 1
    for i := 2; i <= x; i++ {
        f[i] = f[i-1] + f[i-2]
    }
    return f[x]
}