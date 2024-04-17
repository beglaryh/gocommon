package list

import (
	"github.com/beglaryh/gocommon"
	"github.com/beglaryh/gocommon/collection"
)

type List[T any] interface {
	collection.Collection[T]
	Remove(index int) gocommon.Optional[T]
	Get(index int) gocommon.Optional[T]
}
