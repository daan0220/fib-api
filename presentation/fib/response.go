package fib

type Response struct {
	Result string `json:"result"`
}

func NewResponse(res string) Response {
	return Response{
		Result: res,
	}
}
