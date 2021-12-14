package web

type Response struct {
	Code    int         `json:"code"`
	Content interface{} `json:"content, omitempty"`
	Error   string      `json:"error, omitempty"`
}

func NewResponse(code int, content interface{}, err string) Response {
	if code < 300 {
		return Response{code, content, ""}

	}
	return Response{code, nil, err}

}
