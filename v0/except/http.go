package except

import (
	"fmt"
	creq "github.com/Foxcapades/Go-ChainRequest"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/opt"
)

type HttpRequestError interface {
	error

	ResponseCode() *int
	ResponseBody() []byte
}

func NewHttpRequestError(res creq.Response) HttpRequestError {
	if err := res.GetError(); err != nil {
		return &httpReqErr{message: err.Error()}
	} else {
		code := int(res.MustGetResponseCode())
		return &httpReqErr{
			message: fmt.Sprintf("server responded with error code %d", code),
			code:    opt.IntPtr(code),
			body:    res.MustGetBody(),
		}
	}
}

type httpReqErr struct {
	message string
	code    *int
	body    []byte
}

func (h *httpReqErr) Error() string {
	return h.message
}

func (h *httpReqErr) ResponseCode() *int {
	return h.code
}

func (h *httpReqErr) ResponseBody() []byte {
	return h.body
}
