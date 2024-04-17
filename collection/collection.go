package collection

import "errors"

type Collection[T any] interface {
	Add(t ...T) error
	Size() int
	IsEmpty() bool
	Clear()
}

type CollectionError error

var (
	ErrorCollectionLimit = errors.New("collection limit reached. unable to add new element")
)
