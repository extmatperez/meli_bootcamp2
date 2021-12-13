package web

type Response struct {
	Code    int         `json:"code"`
	Content interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func NewReponse(codeStatus int, data interface{}, err string) Response {
	if codeStatus < 300 {
		return Response{codeStatus, data, ""}
	} else {
		return Response{codeStatus, nil, err}
	}
}
