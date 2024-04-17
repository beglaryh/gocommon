package collection

import (
	"github.com/beglaryh/gocommon"
)

type List[T any] interface {
	Collection[T]
	Remove(index int) gocommon.Optional[T]
	Get(index int) gocommon.Optional[T]
}
