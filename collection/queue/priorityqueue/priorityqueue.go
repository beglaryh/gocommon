package priorityqueue

import (
	"errors"

	"github.com/beglaryh/gocommon/collection/list/linkedlist"
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

	if node == nil {
		var t T
		return t, errors.New("queue is empty")
	}
	pq.head = nil
	pq.removeAndReorder(node)
	pq.size -= 1
	return node.GetValue(), nil
}

func (pq *PriorityQueue[T]) Size() int {
	return pq.size
}

func (pq *PriorityQueue[T]) IsEmpty() bool {
	return pq.size == 0
}

func (pq *PriorityQueue[T]) reorder(node *treenode.TreeNode[T]) {
	parent := node.Parent
	for parent != nil && pq.comparator(node.GetValue(), parent.GetValue()) > 0 {
		pq.swap(parent, node)
		parent = node.Parent
	}

}

func (pq *PriorityQueue[T]) removeAndReorder(node *treenode.TreeNode[T]) {
	left := node.Left
	right := node.Right

	for !left.IsEmpty() {
		if !left.IsEmpty() && right.IsEmpty() {
			if pq.comparator(left.GetValue(), right.GetValue()) > 1 {

			} else {

			}

		} else if !left.IsEmpty() {

		} else if !right.IsEmpty() {

		}
	}

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
