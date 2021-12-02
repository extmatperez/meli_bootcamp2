package web

type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func NewResponse(code int, data interface{}, err string) Response {
	if code < 300 {
		return Response{
			Code:  code,
			Data:  data,
			Error: "",
		}
	}

	return Response{
		Code:  code,
		Data:  nil,
		Error: err,
	}
}
