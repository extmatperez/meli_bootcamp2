package models

type Transaction struct {
	ID                  int     `json:"id"`
	CodigoDeTransaccion string  `json:"codigo_de_transaccion" binding:"required"`
	Moneda              string  `json:"moneda" binding:"required"`
	Monto               float64 `json:"monto" binding:"required"`
	Emisor              string  `json:"emisor" binding:"required"`
	Receptor            string  `json:"receptor" binding:"required"`
	FechaDeTransaccion  string  `json:"fecha_de_transaccion" binding:"required"`
}
