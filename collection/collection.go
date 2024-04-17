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
	ERROR_COLLECTION_LIMIT = errors.New("collection limit reached. unable to add new element")
)
