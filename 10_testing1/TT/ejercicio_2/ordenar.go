package ordenar

func Ordenar(vecOr []int) []int {
	vec := vecOr
	for k := 0; k < len(vec)-1; k++ {
		for f := 0; f < len(vec)-1-k; f++ {
			if vec[f] > vec[f+1] {
				aux := vec[f]
				vec[f] = vec[f+1]
				vec[f+1] = aux
			}
		}
	}
	return vec
}
