package treenode

type TreeNode[T any] struct {
	Value  T
	Parent *TreeNode[T]
	Left   *TreeNode[T]
	Right  *TreeNode[T]
}

func New[T any](v T) *TreeNode[T] {
	return &TreeNode[T]{
		Value: v,
	}
}

func (t TreeNode[T]) IsRich() bool {
	return t.Left != nil && t.Right != nil
}
