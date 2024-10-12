package priorityqueue

import (
	"errors"

	"github.com/beglaryh/gocommon/collection/list/linkedlist"
	"github.com/beglaryh/gocommon/treenode"
)

type PriorityQueue[T comparable] struct {
	head       *treenode.TreeNode[T]
	leafs      *linkedlist.LinkedList[*treenode.TreeNode[T]]
	comparator func(T, T) int
	size       int
}

func New[T comparable](comparator func(T, T) int) *PriorityQueue[T] {
	ll := linkedlist.New[*treenode.TreeNode[T]]()
	return &PriorityQueue[T]{
		leafs:      &ll,
		comparator: comparator,
	}
}

func (pq *PriorityQueue[T]) Add(t ...T) error {
	for _, e := range t {
		node := treenode.New[T](e)
		if pq.IsEmpty() {
			pq.head = node
		} else {
			for {
				leaf, _ := pq.leafs.Peek()
				if leaf.IsRich() {
					pq.leafs.Poll()
				} else {
					break
				}
			}
			leaf, _ := pq.leafs.Peek()

			if leaf.Left == nil {
				leaf.Left = node
			} else {
				leaf.Right = node
				pq.leafs.Poll()
			}
			node.Parent = leaf
		}
		pq.leafs.Add(node)
		pq.size += 1
		if pq.size > 1 {
			pq.reorder()
		}
	}

	return nil
}

func (pq *PriorityQueue[T]) Peek() (T, error) {
	if pq.IsEmpty() {
		var t T
		return t, errors.New("queue is empty")
	}
	return pq.head.Value, nil
}

func (pq *PriorityQueue[T]) Poll() (T, error) {
	node := pq.head

	if node == nil {
		var t T
		return t, errors.New("queue is empty")
	}

	pq.removeAndReorder(node)
	return node.Value, nil
}

func (pq *PriorityQueue[T]) Size() int {
	return pq.size
}

func (pq *PriorityQueue[T]) IsEmpty() bool {
	return pq.size == 0
}

func (pq *PriorityQueue[T]) reorder() {
	node, _ := pq.leafs.Get(-1)
	parent := node.Parent
	for parent != nil && pq.comparator(node.Value, parent.Value) > 0 {

		if node != parent.Left {
			node.Left = parent.Left
		} else {
			node.Right = parent.Right
		}
		node.Parent = parent.Parent
		parent.Parent = node
		parent.Left = nil
		parent.Right = nil
		if parent == pq.head {
			pq.head = node
		}
		parent = node.Parent
	}

}

func (pq *PriorityQueue[T]) removeAndReorder(node *treenode.TreeNode[T]) {
	left := node.Left
	right := node.Right
	var selected *treenode.TreeNode[T]

	if left != nil && right != nil {
		comp := pq.comparator(left.Value, right.Value)
		if comp > 0 {
			selected = left
			pq.removeAndReorder(left)
			left.Right = right
			right.Parent = left
		} else {
			selected = right
			pq.removeAndReorder(right)
			right.Left = left
			left.Parent = right
		}
	} else if left != nil {
		selected = left
	} else if right != nil {
		selected = right
	}

	if pq.head == node {
		pq.head = selected
	}
}
