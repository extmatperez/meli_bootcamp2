package turnotarde

import (
	"errors"
	"sort"
)

func Restar(a,b int) int{
	return a-b
}

func Dividir(a,b int) (int,error){
	if(b == 0){
		return 0,errors.New("no se puede dividir por 0")
	}
	return a/b,nil
}

func OrdernarSilice(silice ...int) []int{
	 sort.Ints(silice)
//	desc sort.Sort(sort.Reverse(sort.IntSlice(silice)))
//	fmt.Println(silice)
	return silice
}