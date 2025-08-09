package models

type Response struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data any `json:"data"`
}


func (r *Response) WithCode (code int) *Response {
	r.Code = code
	return r
}

func (r *Response) WithMessage (msg string) *Response {
	r.Message = msg
	return r
}

func (r *Response) WithData (data any) *Response {
	r.Data = data
	return r
}