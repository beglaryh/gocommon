package fifo

import (
	"github.com/beglaryh/gocommon/collection/list/linkedlist"
	"github.com/beglaryh/gocommon/collection/stream"
)

type FifoQueue[T comparable] linkedlist.LinkedList[T]

func NewFifoQueue[T comparable]() *FifoQueue[T] {
	return &FifoQueue[T]{}
}

func (q *FifoQueue[T]) Add(t ...T) error {
	return (*linkedlist.LinkedList[T])(q).Add(t...)
}
func (q *FifoQueue[T]) Peek() (T, error) {
	return (*linkedlist.LinkedList[T])(q).Get(0)
}

func (q *FifoQueue[T]) Remove() (T, error) {
	return (*linkedlist.LinkedList[T])(q).Remove(0)
}

func (q *FifoQueue[T]) Size() int {
	return (*linkedlist.LinkedList[T])(q).Size()
}

func (q *FifoQueue[T]) IsEmpty() bool {
	return q.Size() == 0
}

func (q *FifoQueue[T]) Clear() {
	(*linkedlist.LinkedList[T])(q).Clear()
}

func (q *FifoQueue[T]) Stream() stream.Stream[T] {
	return (*linkedlist.LinkedList[T])(q).Stream()
}
