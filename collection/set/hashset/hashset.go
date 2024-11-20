package hashset

import "github.com/beglaryh/gocommon/stream"

type HashSet[T comparable] struct {
	set map[T]bool
}

func New[T comparable]() *HashSet[T] {
	return &HashSet[T]{
		set: map[T]bool{},
	}
}

func (s *HashSet[T]) Add(t ...T) error {
	for _, e := range t {
		s.add(e)
	}
	return nil
}

func (s *HashSet[T]) add(t T) error {
	s.set[t] = true
	return nil
}

func (s *HashSet[T]) Contains(t T) bool {
	return s.set[t]
}

func (s *HashSet[T]) ToArray() []T {
	l := make([]T, s.Size())
	i := 0
	for k := range s.set {
		l[i] = k
		i += 1
	}
	return l
}

func (s *HashSet[T]) Size() int {
	return len(s.set)
}

func (s *HashSet[T]) IsEmpty() bool {
	return len(s.set) == 0
}

func (s *HashSet[T]) Clear() {
	s.set = map[T]bool{}
}

func (s *HashSet[T]) Stream() *stream.Stream[T] {
	return stream.Of(s.ToArray())
}

func (s *HashSet[T]) Iter(yield func(int, T) bool) {
	i := 0
	for k := range s.set {
		if !yield(i, k) {
			return
		}
		i += 1
	}
}
