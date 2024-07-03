package arraylist

import "github.com/beglaryh/gocommon/collection/collection_errors"

type Builder[T comparable] struct {
	elements []T
	limit    int
}

func NewBuilder[T comparable]() *Builder[T] {
	return &Builder[T]{}
}

func (b *Builder[T]) WithElements(elements []T) *Builder[T] {
	b.elements = elements
	return b
}

func (b *Builder[T]) WithLimit(limit int) *Builder[T] {
	b.limit = limit
	return b
}

func (b *Builder[T]) Build() (*ArrayList[T], error) {
	size := len(b.elements)
	if size > b.limit {
		return nil, collection_errors.LimitExceeded
	}
	initialCapacity := 10
	if size > initialCapacity {
		initialCapacity = size + (size / 2)
	}
	elements := make([]T, initialCapacity)
	copy(elements, b.elements)
	return &ArrayList[T]{
		elements: elements,
		size:     size,
		limit:    b.limit,
	}, nil
}
