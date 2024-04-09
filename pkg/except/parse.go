package except

import creq "github.com/Foxcapades/Go-ChainRequest"

type ParseError interface {
	error

	ResponseCode() int
	ResponseBody() []byte
}

func NewParseError(err error, response creq.Response) ParseError
