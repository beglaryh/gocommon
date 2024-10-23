package set

import "github.com/beglaryh/gocommon/collection"

type Set[T comparable] interface {
	collection.Collection[T]
}
