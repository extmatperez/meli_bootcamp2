package models

type Producto struct{
	ID int
	Nombre string
	Marca string
}

type Transaccion struct {
	ID                int     `json:"id"`
	CodigoTransaccion int     `json:"codigo_transaccion"`
	Moneda            string  `json:"moneda"`
	Monto             float64 `json:"monto"`
	Emisor            string  `json:"emisor"`
	Receptor          string  `json:"receptor"`
	FechaTransaccion  string  `json:"fecha_transaccion"`
	Articulo Producto
}