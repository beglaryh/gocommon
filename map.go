package gocommon

type Map[K comparable, V any] struct {
	m map[K]V
}

func NewMap[K comparable, V any]() *Map[K, V] {
	m := Map[K, V]{}
	return &m
}

func (m *Map[K, V]) Get(k K) Optional[V] {
	v, ok := m.m[k]
	if !ok {
		return Empty[V]()
	}
	return With[V](v)
}
