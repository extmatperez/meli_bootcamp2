package web

type Response struct {
	Code    int         `json:"code"`
	Content interface{} `json:"data, omitempty"`
	Error   string      `json:"error, omitempty"`
}

func NewResponse(statusCode int, data interface{}, err string) Response {
	if statusCode < 300 {
		return Response{statusCode, data, ""}
	} else {
		return Response{statusCode, nil, err}
	}
}
