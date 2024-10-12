package arraylist

import (
	"reflect"

	"github.com/beglaryh/gocommon/collection/collection_errors"
	"github.com/beglaryh/gocommon/collection/stream"
)

type ArrayList[T comparable] struct {
	elements []T
	size     int
	limit    int
}

func New[T comparable]() (*ArrayList[T], error) {
	al := &ArrayList[T]{
		elements: make([]T, 10),
	}
	return al, nil
}

func (l *ArrayList[T]) ToArray() []T {
	return l.elements[0:l.size]
}

func (l *ArrayList[T]) Add(t ...T) error {
	if l.limit != 0 && l.size == l.limit {
		return collection_errors.LimitExceeded
	}
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
	return nil
}

func (l *ArrayList[T]) Remove(index int) (T, error) {
	t, err := l.Get(index)
	if err != nil {
		return t, err
	}
	if index < 0 {
		index = l.size + index
	}

	for i := range l.size - index {
		if (i + index + 1) < l.size {
			l.elements[i+index] = l.elements[i+index+1]
		} else if i == l.size-1 {
			var t T
			l.elements[i] = t
		}
	}
	l.size -= 1
	return t, nil
}

func (l *ArrayList[T]) RemoveValue(t T) int {
	for i, v := range l.elements {
		if v == t {
			_, _ = l.Remove(i)
			return i
		}
	}
	return -1
}

func (l *ArrayList[T]) Get(i int) (T, error) {
	if i < 0 {
		i = l.Size() + i
	}
	if i >= l.Size() {
		var t T
		return t, collection_errors.IndexOutOfBounds
	}

	return l.elements[i], nil
}

func (l *ArrayList[T]) Clear() {
	l.elements = make([]T, 10)
	l.size = 0
}

func (l *ArrayList[T]) Size() int {
	return l.size
}

func (l *ArrayList[T]) IsEmpty() bool {
	return l.Size() == 0
}

func (l *ArrayList[T]) Equals(o *ArrayList[T]) bool {
	return reflect.DeepEqual(l, o)
}

func (l *ArrayList[T]) Stream() stream.Stream[T] {
	return stream.Of[T](l.ToArray())
}

func (l *ArrayList[T]) Iter(yield func(int, T) bool) {
	for i := range l.Size() {
		v, _ := l.Get(i)
		if !yield(i, v) {
			return
		}
	}
}
