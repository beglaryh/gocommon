package queue

import (
	"github.com/beglaryh/gocommon"
	"github.com/beglaryh/gocommon/collection/list"
)

type FifoQueue[T any] list.LinkedList[T]

func NewFifoQueue[T any]() *FifoQueue[T] {
	return &FifoQueue[T]{}
}

func (q *FifoQueue[T]) Add(t ...T) error {
	return (*list.LinkedList[T])(q).Add(t...)
}
func (q *FifoQueue[T]) Peek() gocommon.Optional[T] {
	return (*list.LinkedList[T])(q).Get(0)
}

func (q *FifoQueue[T]) Remove() gocommon.Optional[T] {
	return (*list.LinkedList[T])(q).Remove(0)
}

func (q *FifoQueue[T]) Size() int {
	return (*list.LinkedList[T])(q).Size()
}

func (q *FifoQueue[T]) IsEmpty() bool {
	return q.Size() == 0
}

func (q *FifoQueue[T]) Clear() {
	(*list.LinkedList[T])(q).Clear()
}
