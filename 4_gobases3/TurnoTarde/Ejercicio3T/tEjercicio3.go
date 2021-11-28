package main

import (
	"fmt"

)

type Mantenimiento struct {
	Nombre string
	Precio float64
} 
type Servicio struct {
	Nombre string
	Precio float64
	MinutosTrabajados float64

} 
type Producto struct {
	Nombre string
	Precio float64
	Cantidad int
} 

func SumarProductos(c chan float64,productos... Producto) float64{
	var  precioTotal float64
	z:=0
	fmt.Println("Ejecutando SumarProductos")
	for _,producto := range productos{
		fmt.Println("SumandoProd",z)
		z++
		precioTotal += producto.Precio * float64(producto.Cantidad)
	}
	fmt.Println("Terminado SumarProductos")
	c <- precioTotal
	return precioTotal

}

func SumarMantenimieto(c chan float64,mantinimientos... Mantenimiento) float64{
	var  precioTotal float64
	z:=0
	fmt.Println("Ejecutando Sumar Mantenimiento")
	for _,mantenimiento := range mantinimientos{
		fmt.Println("SumandoMan",z)
		z++
		precioTotal += mantenimiento.Precio 
	}
	fmt.Println("Terminado Sumar Mantenimiento")
	c <- precioTotal
	return precioTotal

}


func SumarServicos(c chan float64,servicios... Servicio) float64{
	var  precioTotal float64
	z:=0
	fmt.Println("Ejecutando SumarServicios")
	for _,servicio := range servicios{
		fmt.Println("SumandoServ",z)
		z++
		if(servicio.MinutosTrabajados > 30){
			precioTotal += servicio.Precio * servicio.MinutosTrabajados
		}else{
			precioTotal += servicio.Precio * 30
		}
		
	}
	fmt.Println("Terminado SumarServicios")
	c <- precioTotal
	return precioTotal

}

func main() {
	

	
	servicio1 := Servicio{Nombre: "Construccion",Precio: 100, MinutosTrabajados: 80}
	servicio2 := Servicio{Nombre: "Repacion baño",Precio: 50, MinutosTrabajados: 50}


	producto1 := Producto{Nombre: "leche",Precio: 50,Cantidad: 4}
	producto2 := Producto{Nombre: "carne",Precio: 500,Cantidad: 4}

	mantimiento1 := Mantenimiento {Nombre: "Casas",Precio: 1000}
	mantimiento2 := Mantenimiento {Nombre: "Edficios",Precio: 2000}
	mantimiento3 := Mantenimiento {Nombre: "Edficios",Precio: 8000}
	mantimiento4 := Mantenimiento {Nombre: "Edficios",Precio: 9000}
	mantimiento5 := Mantenimiento {Nombre: "Edficios",Precio: 10000}
	mantimiento6 := Mantenimiento {Nombre: "Edficios",Precio: 11000}
	mantimiento7 := Mantenimiento {Nombre: "Edficios",Precio: 20200}
	mantimiento8 := Mantenimiento {Nombre: "Edficios",Precio: 200520}

	c := make(chan float64)
	go SumarMantenimieto(c,mantimiento1,mantimiento2,mantimiento3,mantimiento4,mantimiento5,mantimiento6,mantimiento7,mantimiento8)
	go SumarServicos(c,servicio1,servicio2)
	go SumarProductos(c,producto1,producto2)
	go SumarProductos(c,producto1,producto2)
	go SumarProductos(c,producto1,producto2)
	go SumarProductos(c,producto1,producto2)

	fmt.Println(<-c+<-c+<-c+<-c+<-c+<-c)



	

}

