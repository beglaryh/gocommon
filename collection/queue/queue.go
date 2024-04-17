package queue

import (
	"github.com/beglaryh/gocommon"
	"github.com/beglaryh/gocommon/collection"
)

type Queue[T any] interface {
	collection.Collection[T]
	Remove() gocommon.Optional[T]
	Peek() gocommon.Optional[T]
}
