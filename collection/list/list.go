package list

import "github.com/beglaryh/gocommon/collection"

type List[T comparable] interface {
	collection.Collection[T]
	Remove(index int) (T, error)
	Get(index int) (T, error)
	RemoveValue(t T) int
}
