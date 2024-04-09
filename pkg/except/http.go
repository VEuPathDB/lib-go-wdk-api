package except

import creq "github.com/Foxcapades/Go-ChainRequest"

type HttpRequestError interface {
	error

	ResponseCode() *int
	ResponseBody() []byte
}

func NewHttpRequestError(res creq.Response) HttpRequestError
