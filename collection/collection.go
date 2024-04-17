package collection

import "errors"

type Collection[T any] interface {
	Add(t ...T) error
	Size() int
	IsEmpty() bool
	Clear()
	Stream() Stream[T]
}

type Error error

var (
	ErrorCollectionLimit Error = errors.New("collection limit reached. unable to add new element")
)
