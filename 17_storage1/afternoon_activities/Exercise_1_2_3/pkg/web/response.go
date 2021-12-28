package web

// Definimos la estructura de respuesta
type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

// Creamos una función que nos permita crear un response y nos la devuelva
func New_response(code_status int, data interface{}, err string) Response {
	if code_status < 300 && code_status != 201 {
		return Response{code_status, data, ""}
	} else {
		return Response{code_status, nil, err}
	}
}
