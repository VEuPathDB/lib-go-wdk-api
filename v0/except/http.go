package except

import (
	"fmt"
	creq "github.com/Foxcapades/Go-ChainRequest"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/opt"
)

type HttpRequestError interface {
	error

	ResponseCode() opt.Int
	ResponseBody() []byte
}

func NewHttpRequestError(res creq.Response) HttpRequestError {
	if err := res.GetError(); err != nil {
		return &httpReqErr{message: err.Error()}
	} else {
		code := int(res.MustGetResponseCode())
		return &httpReqErr{
			message: fmt.Sprintf("server responded with error code %d", code),
			code:    opt.NewInt(code),
			body:    res.MustGetBody(),
		}
	}
}

type httpReqErr struct {
	message string
	code    opt.Int
	body    []byte
}

func (h *httpReqErr) Error() string {
	return h.message
}

func (h *httpReqErr) ResponseCode() opt.Int {
	return h.code
}

func (h *httpReqErr) ResponseBody() []byte {
	return h.body
}
