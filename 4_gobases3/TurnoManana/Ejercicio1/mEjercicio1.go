package main

import (
	"encoding/json"
	"fmt"
	"os"
	
)

type Product struct {
	Id int `json:"id"`
	Precio float64 `json:"precio"`
	Cantidad int `json:"cantidad"`
}



func main() {
	Product1 := Product{Id: 1,Precio: 15.5,Cantidad: 10}
	Product2 := Product{Id: 2,Precio: 25.5,Cantidad: 8}


	 ListProductos := [] Product{Product1,Product2}
	 encjson, _ := json.Marshal(ListProductos)
	ruta := "./archivo.csv"
	 ruta, err :=SaveDocument(encjson,ruta)
	if(err != nil){

		fmt.Printf("Ocurrio un error %w",err)
	}else{
		fmt.Printf("El archivo %s se guardo con exito",ruta)
	}

	ReadDocument(ruta)

 }


func SaveDocument (archivo []byte, ruta string) (string,error){
	
	err := os.WriteFile(ruta, archivo,0644)

	if(err != nil) {
		return "",err
	}

	return ruta,nil

}

func ReadDocument (ruta string) (error){

	data,err := os.ReadFile(ruta)
	var ListProductos  []Product
	if(err != nil){
		return err
	}
		json.Unmarshal(data, &ListProductos)
		fmt.Printf("\n%v\t%v\t%v", "Id", "Precio", "Cantidad")
		for _,product := range ListProductos{
			fmt.Printf("\n")
			
			fmt.Printf("%v,\t%.2f,\t%v",product.Id,product.Precio,product.Cantidad)
			

		}
		return nil
	

}
