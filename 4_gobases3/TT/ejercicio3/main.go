package main

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

type Servicio struct {
	Nombre            string
	Precio            float64
	MinutosTrabajados int
}

type Mantenimiento struct {
	Nombre string
	Precio float64
}

func SumarProductos(prods []Producto) float64 {
	acum := 0.0
	for _, p := range prods {
		acum += p.Precio * float64(p.Cantidad)
	}
	return acum
}

func SumarServicios(servs []Servicio) float64 {
	acum := 0.0
	for _, s := range servs {
		if s.MinutosTrabajados < 30 {
			acum += s.Precio
		} else {
			acum += s.Precio * (float64(s.MinutosTrabajados / 30))
		}
	}
	return acum
}

func SumarMantenimiento(mant []Mantenimiento) float64 {
	acum := 0.0
	for _, m := range mant {
		acum += m.Precio
	}
	return acum
}

// var listaProd []Producto = [{"lentes", 30.5, 2}, ]
// listaService := []
// listaMantenimiento := []
