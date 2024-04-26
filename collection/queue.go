package collection

type Queue[T any] interface {
	Collection[T]
	Remove() (T, error)
	Peek() (T, error)
}
