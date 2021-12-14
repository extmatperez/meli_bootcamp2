package web

type response struct {
	Code      int         `json:"code"`
	Contenido interface{} `json:"data,omitempty"`
	Error     string      `json:"error,omitempty"`
}

func NewResponse(codeStatus int, data interface{}, err string) response {
	if codeStatus < 300 {
		return response{codeStatus, data, ""}
	} else {
		return response{codeStatus, nil, err}
	}
}
