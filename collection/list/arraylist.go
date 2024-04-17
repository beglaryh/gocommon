package list

import (
	"errors"
	"github.com/beglaryh/gocommon"
	"github.com/beglaryh/gocommon/collection"
	"reflect"
)

type ArrayList[T any] struct {
	elements []T
	size     int
	limit    int
}

type ArrayListOption[T any] func(*ArrayList[T]) error

func ArrayListWithLimit[T any](limit int) ArrayListOption[T] {
	return func(l *ArrayList[T]) error {
		if limit < 1 {
			return errors.New("invalid limit. expecting greater than 0")

		}
		l.limit = limit
		return nil
	}
}

func ArrayListWithInitialCapacity[T any](capacity int) ArrayListOption[T] {
	return func(l *ArrayList[T]) error {
		if capacity < 1 {
			return errors.New("invalid initial capacity. expecting greater than 0")
		}
		l.elements = make([]T, capacity)
		return nil
	}
}

func ArrayListWithSlice[T any](slice []T) ArrayListOption[T] {
	return func(l *ArrayList[T]) error {
		if l.limit != 0 && l.limit < len(slice) {
			return errors.New("list limit too small for given slice")
		}
		l.elements = make([]T, len(slice)+(len(slice)/2))
		copy(l.elements, slice)
		l.size = len(slice)
		return nil
	}
}

func NewArrayList[T any](options ...ArrayListOption[T]) (*ArrayList[T], error) {
	al := &ArrayList[T]{
		elements: make([]T, 10),
	}

	for _, option := range options {
		err := option(al)
		if err != nil {
			return nil, err
		}
	}
	return al, nil
}

func (l *ArrayList[T]) ToArray() []T {
	return l.elements[0:l.size]
}

func (l *ArrayList[T]) Add(t ...T) error {
	if l.limit != 0 && l.size == l.limit {
		return collection.ErrorCollectionLimit
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

func (l *ArrayList[T]) Remove(index int) gocommon.Optional[T] {
	return gocommon.Empty[T]()
}

func (l *ArrayList[T]) Get(i int) gocommon.Optional[T] {
	if i < 0 {
		i = l.Size() + i
	}
	if i >= l.Size() {
		return gocommon.Empty[T]()
	}

	return gocommon.WithPointer[T](&l.elements[i])
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
