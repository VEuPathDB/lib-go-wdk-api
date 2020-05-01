package common

type PostBody interface {
	Validate() error
}
