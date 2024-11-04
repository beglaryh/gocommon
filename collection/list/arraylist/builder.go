package arraylist

import "github.com/beglaryh/gocommon/collection/collection_errors"

type Builder[T comparable] struct {
	elements        []T
	initialCapacity int
	limit           int
}

func NewBuilder[T comparable]() *Builder[T] {
	return &Builder[T]{}
}

func (b *Builder[T]) WithElements(elements []T) *Builder[T] {
	b.elements = elements
	return b
}

func (b *Builder[T]) WithInitialCapacity(capacity int) *Builder[T] {
	b.initialCapacity = capacity
	return b
}

func (b *Builder[T]) WithLimit(limit int) *Builder[T] {
	b.limit = limit
	return b
}

func (b *Builder[T]) Build() (*ArrayList[T], error) {
	size := len(b.elements)
	if b.limit != 0 && size > b.limit {
		return nil, collection_errors.LimitExceeded
	}
	if b.initialCapacity <= 0 {
		b.initialCapacity = 10
	}
	initialCapacity := b.initialCapacity
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
