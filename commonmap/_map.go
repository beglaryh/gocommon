package commonmap

import "github.com/beglaryh/gocommon"

type Map[K comparable, V any] interface {
	Get(k K) gocommon.Optional[V]
	GetOrDefault(k K, defaultValue V) V
	Put(k K, v V) gocommon.Optional[V]
	Values() []V
	Contains(k K) bool
}
