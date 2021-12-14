package web

type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func NewResponse(statusCode int, data interface{}, error string) Response {
	if statusCode < 300 {
		return Response{statusCode, data, ""}
	} else {
		return Response{statusCode, nil, error}
	}

}
