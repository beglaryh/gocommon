package linkedhashmap

import (
	"github.com/beglaryh/gocommon/collection/list/linkedlist"
	"github.com/beglaryh/gocommon/commonmap/hashmap"
	"github.com/beglaryh/gocommon/optional"
)

type LinkedHashMap[K comparable, V any] struct {
	m    hashmap.HashMap[K, V]
	keys linkedlist.LinkedList[K]
}

func New[K comparable, V any]() LinkedHashMap[K, V] {
	return LinkedHashMap[K, V]{
		keys: linkedlist.New[K](),
	}
}

func (m *LinkedHashMap[K, V]) Get(k K) optional.Optional[V] {
	return m.m.Get(k)
}

func (m *LinkedHashMap[K, V]) GetOrDefault(k K, v V) V {
	return m.m.GetOrDefault(k, v)
}

func (m *LinkedHashMap[K, V]) Put(k K, v V) optional.Optional[V] {
	op := m.m.Put(k, v)
	if op.IsPresent() {
		m.keys.RemoveValue(k)
	}
	_ = m.keys.Add(k)
	return op
}

func (m *LinkedHashMap[K, V]) Remove(k K) optional.Optional[V] {
	vop := m.m.Remove(k)
	if vop.IsPresent() {
		m.keys.RemoveValue(k)
	}
	return vop
}

func (m *LinkedHashMap[K, V]) Contains(k K) bool {
	return m.m.Contains(k)
}

func (m *LinkedHashMap[K, V]) Compute(k K, compute func(k K, v optional.Optional[V]) optional.Optional[V]) optional.Optional[V] {
	return m.m.Compute(k, compute)
}

func (m *LinkedHashMap[K, V]) ComputeIfAbsent(k K, compute func(k K, v optional.Optional[V]) optional.Optional[V]) optional.Optional[V] {
	return m.m.ComputeIfAbsent(k, compute)
}

func (m *LinkedHashMap[K, V]) Size() int {
	return m.m.Size()
}

func (m *LinkedHashMap[K, V]) Values() []V {
	values := make([]V, m.Size())
	for i, k := range m.keys.ToArray() {
		v, _ := m.Get(k).Get()
		values[i] = v
	}
	return values
}
