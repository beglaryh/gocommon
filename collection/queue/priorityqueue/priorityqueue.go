package priorityqueue

import (
	"errors"

	"github.com/beglaryh/gocommon/collection/list/linkedlist"
	"github.com/beglaryh/gocommon/stream"
	"github.com/beglaryh/gocommon/treenode"
)

type PriorityQueue[T comparable] struct {
	head       *treenode.TreeNode[T]
	comparator func(T, T) int
	size       int
	vacancies  linkedlist.LinkedList[*treenode.TreeNode[T]]
}

func New[T comparable](comparator func(T, T) int) *PriorityQueue[T] {
	vacancies := linkedlist.New[*treenode.TreeNode[T]]()
	var t T
	ve := treenode.New[T](t)
	vacancies.Add(ve)
	return &PriorityQueue[T]{
		comparator: comparator,
		vacancies:  vacancies,
		head:       ve,
	}
}

func (pq *PriorityQueue[T]) Add(t ...T) error {
	for _, e := range t {

		node, _ := pq.vacancies.Remove(0)
		node.SetValue(e)
		left := treenode.Empty[T]()
		right := treenode.Empty[T]()
		left.Parent = node
		right.Parent = node
		node.Left = left
		node.Right = right
		pq.vacancies.Add(left, right)

		pq.size += 1
		if pq.size > 1 {
			pq.reorder(node)
		}
	}

	return nil
}

func (pq *PriorityQueue[T]) Peek() (T, error) {
	if pq.IsEmpty() {
		var t T
		return t, errors.New("queue is empty")
	}
	return pq.head.GetValue(), nil
}

func (pq *PriorityQueue[T]) Poll() (T, error) {
	node := pq.head

	if pq.IsEmpty() {
		var t T
		return t, errors.New("queue is empty")
	}

	pq.size -= 1
	_ = pq.replace(node)

	return node.GetValue(), nil
}

func (pq *PriorityQueue[T]) Size() int {
	return pq.size
}

func (pq *PriorityQueue[T]) IsEmpty() bool {
	return pq.size == 0
}

func (pq *PriorityQueue[T]) Clear() {
	nq := New[T](pq.comparator)
	pq.head = nq.head
	pq.vacancies = nq.vacancies
}

func (pq *PriorityQueue[T]) ToArray() []T {
	arr := make([]T, pq.size)
	queue := linkedlist.New[*treenode.TreeNode[T]]()
	queue.Add(pq.head)
	i := 0
	for !queue.IsEmpty() {
		e, _ := queue.Poll()
		arr[i] = e.GetValue()
		if !e.Left.IsEmpty() {
			queue.Add(e.Left)
		}
		if !e.Right.IsEmpty() {
			queue.Add(e.Right)
		}
		i += 1
	}
	return arr
}

func (pq *PriorityQueue[T]) Stream() *stream.Stream[T] {
	return stream.Of(pq.ToArray())
}

func (pq *PriorityQueue[T]) Iter(yield func(int, T) bool) {
	queue := linkedlist.New[*treenode.TreeNode[T]]()
	queue.Add(pq.head)
	i := 0
	for !queue.IsEmpty() {
		e, _ := queue.Poll()
		if !yield(i, e.GetValue()) {
			return
		}
		if !e.Left.IsEmpty() {
			queue.Add(e.Left)
		}
		if !e.Right.IsEmpty() {
			queue.Add(e.Right)
		}
		i += 1
	}
}

func (pq *PriorityQueue[T]) reorder(node *treenode.TreeNode[T]) {
	parent := node.Parent
	for parent != nil && pq.comparator(node.GetValue(), parent.GetValue()) > 0 {
		pq.swap(parent, node)
		parent = node.Parent
	}
}

func (pq *PriorityQueue[T]) replace(node *treenode.TreeNode[T]) *treenode.TreeNode[T] {
	if pq.IsEmpty() {
		pq.head = treenode.Empty[T]()
		pq.vacancies.Clear()
		pq.vacancies.Add(pq.head)
		return nil
	}

	left := node.Left
	right := node.Right

	var replacement *treenode.TreeNode[T]
	if !left.IsEmpty() && !right.IsEmpty() {
		replacement = left
		if pq.comparator(left.GetValue(), right.GetValue()) < 0 {
			replacement = right
		}

	} else if !left.IsEmpty() {
		replacement = node.Left
	} else if !right.IsEmpty() {
		replacement = node.Right
	}

	if replacement == nil {
		pq.vacancies.RemoveValue(node.Right)
		return node.Left
	}

	child := pq.replace(replacement)

	if replacement == right {
		left.Parent = replacement
		replacement.Left = left
		replacement.Right = child
	} else {
		right.Parent = replacement
		replacement.Right = right
		replacement.Left = child
	}

	if node == pq.head {
		pq.head = replacement
	}

	return replacement
}

func (pq *PriorityQueue[T]) swap(parent, child *treenode.TreeNode[T]) {
	leftGrandchild := child.Left
	rightGrandchild := child.Right
	grandparent := parent.Parent

	if child == parent.Left {
		child.Left = parent
		child.Right = parent.Right
		parent.Right.Parent = child
	} else {
		child.Right = parent
		child.Left = parent.Left
		parent.Left.Parent = child
	}

	child.Parent = grandparent
	parent.Parent = child

	if grandparent != nil {
		if grandparent.Left == parent {
			grandparent.Left = child
		} else {
			grandparent.Right = child
		}
	}
	parent.Left = leftGrandchild
	parent.Right = rightGrandchild

	leftGrandchild.Parent = parent
	rightGrandchild.Parent = parent

	if parent == pq.head {
		pq.head = child
	}
}
