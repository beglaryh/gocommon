package collection

type List[T any] interface {
	Collection[T]
	Remove(index int) (T, error)
	Get(index int) (T, error)
}
