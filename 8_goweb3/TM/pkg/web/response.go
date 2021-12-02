package web

type response struct {
	Code    int         `json:"code"`
	Content interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// La salida es un response, y el ultimo elemento parametro de la entrada es un string de error.
func NewResponse(code int, data interface{}, err string) response {
	if code < 300 {
		return response{code, data, ""}
	} else {
		return response{code, data, err}
	}
}
