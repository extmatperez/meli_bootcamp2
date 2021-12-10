package ordenamiento

func Ordenar(slice []int) []int {
	var auxiliar int
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice); j++ {
			if slice[i] < slice[j] {
				auxiliar = slice[i]
				slice[i] = slice[j]
				slice[j] = auxiliar
			}
		}
	}
	return slice
}
