package except

import (
	"github.com/Foxcapades/Go-ChainRequest"
)

type ParseError interface {
	error

	ResponseCode() int
	ResponseBody() []byte
}

func NewParseError(err error, response creq.Response) ParseError {
	return &parseError{
		raw:  err,
		code: int(response.MustGetResponseCode()),
		body: response.MustGetBody(),
	}
}

type parseError struct {
	raw  error
	code int
	body []byte
}

func (p *parseError) Error() string {
	return p.raw.Error()
}

func (p *parseError) ResponseCode() int {
	return p.code
}

func (p *parseError) ResponseBody() []byte {
	return p.body
}
