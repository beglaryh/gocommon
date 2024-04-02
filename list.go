package gocommon

import (
	"reflect"
)

type List[T any] struct {
	elements []T
	size     int
}

func EmptyList[T any]() *List[T] {
	l := List[T]{elements: make([]T, 10)}
	return &l
}

func (l *List[T]) ToArray() []T {
	return l.elements[0:l.size]
}

func (l *List[T]) Add(t ...T) {
	offset := l.size
	for _, e := range t {
		if len(l.elements) == l.size {
			capacity := l.size + (l.size / 2)
			newElements := make([]T, capacity)
			copy(newElements, l.elements)
			l.elements = newElements
		}
		l.elements[offset] = e
		l.size += 1
		offset += 1
	}
}

func (l *List[T]) Get(i int) Optional[T] {
	if i < 0 {
		i = l.Size() + i
	}
	if i >= l.Size() {
		return Empty[T]()
	}

	return With[T](&l.elements[i])
}

func (l *List[T]) Clear() {
	l.elements = make([]T, 10)
	l.size = 0
}

func (l *List[T]) Size() int {
	return l.size
}

func (l *List[T]) Equals(o *List[T]) bool {
	return reflect.DeepEqual(l, o)
}
