package web

//Este archivo corresponde a la estructura de la respuestas de peticiones web de

type Response struct {
	Code  int
	Data  interface{}
	Error string
}

func NewResponse(codeStatus int, data interface{}, err string) Response {
	//Respuesta correcta -> codigo menor a 3000
	if codeStatus < 300 {
		return Response{codeStatus, data, ""}
	}
	return Response{codeStatus, nil, err}
}
