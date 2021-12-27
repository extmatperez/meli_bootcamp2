package web

type Response struct {
	Code    int         `json:"code"`
	Content interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// La salida es un response, y el ultimo elemento parametro de la entrada es un string de error.
func NewResponse(code int, data interface{}, err string) Response {
	if code < 300 {
		return Response{code, data, ""}
	} else {
		return Response{code, data, err}
	}
}
