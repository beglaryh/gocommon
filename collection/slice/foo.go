package slice

import (
	"github.com/beglaryh/gocommon/collection/list"
	"github.com/beglaryh/gocommon/collection/list/arraylist"
)

type Slice[T comparable] []T

func (f Slice[T]) ToList() list.List[T] {
	al, _ := arraylist.NewBuilder[T]().WithElements(f).Build()
	return list.List[T](al)
}
