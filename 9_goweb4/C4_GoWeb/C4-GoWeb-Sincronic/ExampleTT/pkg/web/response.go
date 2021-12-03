package web

type Response struct {
	Code  int
	Data  interface{}
	Error string
}

func NewResponse(codeStatus int, data interface{}, err string) Response {
	if codeStatus < 300 {
		return Response{codeStatus, data, ""}
	} else {
		return Response{codeStatus, data, err}
	}
}
