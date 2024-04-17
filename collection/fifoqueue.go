package collection

import (
	"github.com/beglaryh/gocommon"
)

type FifoQueue[T any] LinkedList[T]

func NewFifoQueue[T any]() *FifoQueue[T] {
	return &FifoQueue[T]{}
}

func (q *FifoQueue[T]) Add(t ...T) error {
	return (*LinkedList[T])(q).Add(t...)
}
func (q *FifoQueue[T]) Peek() gocommon.Optional[T] {
	return (*LinkedList[T])(q).Get(0)
}

func (q *FifoQueue[T]) Remove() gocommon.Optional[T] {
	return (*LinkedList[T])(q).Remove(0)
}

func (q *FifoQueue[T]) Size() int {
	return (*LinkedList[T])(q).Size()
}

func (q *FifoQueue[T]) IsEmpty() bool {
	return q.Size() == 0
}

func (q *FifoQueue[T]) Clear() {
	(*LinkedList[T])(q).Clear()
}

func (q *FifoQueue[T]) Stream() Stream[T] {
	return (*LinkedList[T])(q).Stream()
}
