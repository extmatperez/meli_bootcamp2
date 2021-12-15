package fino

func Finobacci(n int)int {
	if (n>1){
		return Finobacci(n-1) + Finobacci(n-2);  //función recursiva
	 }	 else if (n==1) {  // caso base
		 return 1;
	 }	 else if (n==0){  // caso base
		 return 0;
	 } else{ //error
		 return -1; 
	 }
}


func GetSerie(tamano int) []int{
	var resultado []int
    for i := 0; i < tamano; i++ {
      resultado = append(resultado,Finobacci(i))
    }
return resultado
}