package collection

import (
	"github.com/beglaryh/gocommon"
)

type node[T any] struct {
	value T
	next  *node[T]
	prev  *node[T]
}

type LinkedList[T any] struct {
	head  *node[T]
	tail  *node[T]
	size  int
	limit int
}

func NewLinkedList[T any]() LinkedList[T] {
	return LinkedList[T]{
		size: 0,
	}
}

func (l *LinkedList[T]) Add(t ...T) error {
	numberOfElements := len(t)
	if l.limit != 0 && l.size+numberOfElements > l.limit {
		return ErrorCollectionLimit
	}
	for _, e := range t {
		newNode := node[T]{value: e}
		if l.size == 0 {
			l.head = &newNode
			l.tail = &newNode
		} else {
			l.tail.next = &newNode
			newNode.prev = l.tail
			l.tail = &newNode
		}
		l.size += 1
	}
	return nil
}

func (l *LinkedList[T]) get(index int) gocommon.Optional[node[T]] {
	if index < 0 {
		index = l.size + index
	}
	if index >= l.size || index < 0 {
		return gocommon.Empty[node[T]]()
	}

	x := float32(index) / float32(l.size)
	var element *node[T]
	if x < 0.5 {
		element = l.head
		for _ = range index - 1 {
			element = element.next
		}
	} else {
		index = l.size - index - 1
		element = l.tail
		for _ = range index {
			element = element.prev

		}
	}
	return gocommon.WithPointer[node[T]](element)
}

func (l *LinkedList[T]) Get(index int) gocommon.Optional[T] {
	n := l.get(index)
	if !n.IsPresent() {
		return gocommon.Empty[T]()
	}
	element, _ := n.Get()
	return gocommon.With[T](element.value)
}

func (l *LinkedList[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

func (l *LinkedList[T]) Size() int {
	return l.size
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.Size() == 0
}

func (l *LinkedList[T]) Peek() gocommon.Optional[T] {
	return l.Get(0)
}

func (l *LinkedList[T]) Remove(index int) gocommon.Optional[T] {

	optional := l.get(index)
	if !optional.IsPresent() {
		return gocommon.Empty[T]()
	}

	element, _ := optional.GetPointer()
	prev := element.prev
	next := element.next

	if prev == nil {
		l.head = next
	} else {
		prev.next = next
	}

	if next == nil {
		l.tail = prev
	} else {
		next.prev = prev
	}
	l.size -= 1

	return gocommon.With[T](element.value)
}

func (l *LinkedList[T]) ToArray() []T {
	arr := make([]T, l.Size())
	e := l.head
	i := 0
	for e != nil {
		v := e.value
		arr[i] = v
		i += 1
		e = e.next
	}
	return arr
}

func (l *LinkedList[T]) Stream() Stream[T] {
	return StreamOf[T](l.ToArray())
}
