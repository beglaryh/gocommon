package collection

import "github.com/beglaryh/gocommon/stream"

type Collection[T comparable] interface {
	Add(t ...T) error
	Size() int
	IsEmpty() bool
	Clear()
	Stream() stream.Stream[T]
	ToArray() []T
	Iter(yield func(int, T) bool)
}
