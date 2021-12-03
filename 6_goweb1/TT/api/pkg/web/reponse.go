package web

type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func NewResponse(codeStatus int, data interface{}, error string) Response {
	if codeStatus < 300 {
		return Response{codeStatus, data, ""}
	} else {
		return Response{codeStatus, nil, error}
	}
}
