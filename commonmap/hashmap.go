package commonmap

import "github.com/beglaryh/gocommon"

type HashMap[K comparable, V any] map[K]V

func NewMap[K comparable, V any]() *HashMap[K, V] {
	m := HashMap[K, V]{}
	return &m
}

func (m *HashMap[K, V]) Get(k K) gocommon.Optional[V] {
	v, ok := (map[K]V)(*m)[k]
	if !ok {
		return gocommon.Empty[V]()
	}
	return gocommon.With[V](v)
}

func (m *HashMap[K, V]) Put(k K, v V) gocommon.Optional[V] {
	prev := m.Get(k)
	(map[K]V)(*m)[k] = v
	return prev
}

func (m *HashMap[K, V]) Contains(k K) bool {
	_, ok := (map[K]V)(*m)[k]
	return ok
}

func (m *HashMap[K, V]) GetOrDefault(k K, defaultValue V) V {
	vop := m.Get(k)
	if vop.IsEmpty() {
		return defaultValue
	}
	v, _ := vop.Get()
	return v
}

func (m *HashMap[K, V]) Size() int {
	return len((map[K]V)(*m))
}

func (m *HashMap[K, V]) Values() []V {
	elements := make([]V, m.Size())
	mm := (map[K]V)(*m)
	i := 0
	for _, v := range mm {
		elements[i] = v
		i += 1
	}
	return elements
}
