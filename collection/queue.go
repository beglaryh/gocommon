package collection

import (
	"github.com/beglaryh/gocommon"
)

type Queue[T any] interface {
	Collection[T]
	Remove() gocommon.Optional[T]
	Peek() gocommon.Optional[T]
}
