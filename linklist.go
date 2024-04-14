package gocommon

type node[T any] struct {
	value *T
	next  *node[T]
	prev  *node[T]
}

type LinkedList[T any] struct {
	head *node[T]
	tail *node[T]
	size int
}

func NewLinkedList[T any]() LinkedList[T] {
	return LinkedList[T]{
		size: 0,
	}
}

func (l LinkedList[T]) Add(t *T) {
	newNode := node[T]{value: t}
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

func (l LinkedList[T]) get(index int) Optional[node[T]] {
	if index < 0 {
		index = l.size + index
	}
	if index >= l.size || index < 0 {
		return Empty[node[T]]()
	}

	x := float32(index) / float32(l.size)
	var element *node[T]
	if x < 0.5 {
		for i := range index {
			if i == 0 {
				element = l.head
			} else {
				element = element.next
			}
		}
	} else {
		index = l.size - index
		for i := range index {
			if i == 0 {
				element = l.tail
			} else {
				element = element.prev
			}
		}
	}
	return With[node[T]](element)
}

func (l LinkedList[T]) Get(index int) Optional[T] {
	n := l.get(index)
	if !n.IsPresent() {
		return Empty[T]()
	}
	return With[T](n.value.value)
}
func (l LinkedList[T]) Remove(index int) Optional[T] {

	optional := l.get(index)
	if !optional.IsPresent() {
		return Empty[T]()
	}

	element := optional.Get()
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

	return With[T](element.value)
}
