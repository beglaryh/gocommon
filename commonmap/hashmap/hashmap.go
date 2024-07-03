package hashmap

import (
	"github.com/beglaryh/gocommon/optional"
)

type HashMap[K comparable, V any] map[K]V

func New[K comparable, V any]() *HashMap[K, V] {
	m := HashMap[K, V]{}
	return &m
}

func (m *HashMap[K, V]) Get(k K) optional.Optional[V] {
	v, ok := (map[K]V)(*m)[k]
	if !ok {
		return optional.Empty[V]()
	}
	return optional.With[V](v)
}

func (m *HashMap[K, V]) Put(k K, v V) optional.Optional[V] {
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

func (m *HashMap[K, V]) Remove(k K) optional.Optional[V] {
	vop := m.Get(k)
	if vop.IsPresent() {
		delete(map[K]V(*m), k)
	}
	return vop
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

func (m *HashMap[K, V]) Compute(k K, compute func(k K, vop optional.Optional[V]) optional.Optional[V]) optional.Optional[V] {
	vop := m.Get(k)
	computed := compute(k, vop)
	computed.IfPresentOrElse(func(v V) { m.Put(k, v) },
		func() { m.Remove(k) },
	)
	return computed
}

func (m *HashMap[K, V]) ComputeIfAbsent(k K, compute func(k K, vop optional.Optional[V]) optional.Optional[V]) optional.Optional[V] {
	vop := m.Get(k)
	if vop.IsEmpty() {
		return compute(k, vop)
	}
	return vop
}
