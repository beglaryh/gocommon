package treenode

type TreeNode[T any] struct {
	value  T
	Parent *TreeNode[T]
	Left   *TreeNode[T]
	Right  *TreeNode[T]
	empty  bool
}

func New[T any](v T) *TreeNode[T] {
	return &TreeNode[T]{
		value: v,
	}
}

func Empty[T any]() *TreeNode[T] {
	return &TreeNode[T]{
		empty: true,
	}
}

func (t TreeNode[T]) IsRich() bool {
	return t.Left != nil && t.Right != nil
}

func (tn *TreeNode[T]) SetValue(t T) {
	tn.value = t
	tn.empty = false
}

func (tn *TreeNode[T]) GetValue() T {
	return tn.value
}

func (tn *TreeNode[T]) IsEmpty() bool {
	return tn.empty
}
