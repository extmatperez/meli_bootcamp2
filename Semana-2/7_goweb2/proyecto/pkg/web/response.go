package web

type Response struct {
	Code      int         `json:"code"`
	Contenido interface{} `json:"data,omitempty"`
	Error     string      `json:"error,omitempty"`
}

func NewResponse(code int, data interface{}, err string) Response {
	if code < 300 {
		return Response{Code: code, Contenido: data}
	}
	return Response{Code: code, Contenido: nil, Error: err}
}
