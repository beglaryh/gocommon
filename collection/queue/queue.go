package queue

import "github.com/beglaryh/gocommon/collection"

type Queue[T comparable] interface {
	collection.Collection[T]
	Poll() (T, error)
	Peek() (T, error)
}
