package gocommon

type Set[T comparable] struct {
	set map[T]bool
}

func (s *Set[T]) Add(t T) bool {
	preexisting := s.set[t]
	s.set[t] = true
	return preexisting
}
func (s *Set[T]) Contains(t T) bool {
	return s.set[t]
}

func (s *Set[T]) ToArray() []T {
	l := make([]T, s.Size())
	i := 0
	for k := range s.set {
		l[i] = k
		i += 1
	}
	return l
}

func (s *Set[T]) Size() int {
	return len(s.set)
}
