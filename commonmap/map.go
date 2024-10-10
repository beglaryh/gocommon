package commonmap

import (
	"github.com/beglaryh/gocommon/optional"
)

type Map[K comparable, V any] interface {
	Get(k K) optional.Optional[V]
	GetOrDefault(k K, defaultValue V) V
	Put(k K, v V) optional.Optional[V]
	Remove(k K) optional.Optional[V]
	Values() []V
	Contains(k K) bool
	Compute(k K, f func(k K, v optional.Optional[V]) optional.Optional[V]) optional.Optional[V]
	ComputeIfAbsent(k K, f func(k K, v optional.Optional[V]) optional.Optional[V]) optional.Optional[V]
}
