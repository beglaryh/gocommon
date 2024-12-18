package stream

import (
	"github.com/beglaryh/gocommon/optional"
)

type Stream[T any] struct {
	ts []T
}

func Of[T comparable](ts []T) *Stream[T] {
	return &Stream[T]{ts: ts}
}

func (stream *Stream[T]) Filter(filter func(t T) bool) *Stream[T] {
	ns := Stream[T]{}
	for _, t := range stream.ts {
		if filter(t) {
			ns.ts = append(ns.ts, t)
		}
	}
	return &ns
}

func Map[F, T any](fs []F, mapper func(f F) T) *Stream[T] {
	ns := Stream[T]{
		ts: make([]T, len(fs)),
	}
	for i, e := range fs {
		ns.ts[i] = mapper(e)
	}
	return &ns
}

func FlatMap[T any](input [][]T) *Stream[T] {
	totalLength := len(input) * len(input[0])
	ts := make([]T, totalLength)

	counter := 0
	for _, array := range input {
		for _, e := range array {
			ts[counter] = e
			counter += 1
		}
	}
	return &Stream[T]{ts: ts}
}

func GroupBy[K comparable, T any](ts []T, getKey func(t T) K) map[K][]T {
	response := map[K][]T{}
	for _, t := range ts {
		key := getKey(t)
		response[key] = append(response[key], t)
	}
	return response
}

func (stream *Stream[T]) Peek(peekFunc func(t T)) *Stream[T] {
	go peek(stream, peekFunc)
	return stream
}

func (stream *Stream[T]) Slice() []T {
	return stream.ts
}

func peek[T any](s *Stream[T], peekFunc func(t T)) {
	for _, t := range s.ts {
		peekFunc(t)
	}
}

func (stream *Stream[T]) Reduce(identity T, reduce func(a, b T) T) T {
	var a T
	for i, e := range stream.ts {
		if i == 0 {
			a = reduce(identity, e)
		} else {
			a = reduce(a, e)
		}
	}
	return a
}

func (stream *Stream[T]) Sort(sortFunction func(a, b T) bool) *Stream[T] {
	ns := mergeSort(stream.ts, sortFunction)
	return &Stream[T]{ts: ns}
}

func (stream *Stream[T]) AnyMatch(anyFunction func(t T) bool) bool {
	for _, e := range stream.ts {
		if anyFunction(e) {
			return true
		}
	}
	return false
}

func (stream *Stream[T]) NoneMatch(anyFunction func(t T) bool) bool {
	for _, e := range stream.ts {
		if anyFunction(e) {
			return false
		}
	}
	return true
}

func (stream *Stream[T]) FindFirst() optional.Optional[T] {
	if len(stream.ts) == 0 {
		return optional.Empty[T]()
	}
	return optional.With(stream.ts[0])
}

func (stream *Stream[T]) ForEach(forEach func(t T)) {
	for _, t := range stream.ts {
		forEach(t)
	}
}

func mergeSort[T any](es []T, compare func(a, b T) bool) []T {
	if len(es) < 2 {
		return es
	}
	mid := len(es) / 2
	left := mergeSort(es[:mid], compare)
	right := mergeSort(es[mid:], compare)

	var sorted []T
	i, j := 0, 0
	for len(sorted) != len(es) {
		if i == len(left) {
			sorted = append(sorted, right[j])
			j += 1
		} else if j == len(right) {
			sorted = append(sorted, left[i])
			i += 1
		} else {
			lv := left[i]
			rv := right[j]
			if compare(lv, rv) {
				sorted = append(sorted, lv)
				i += 1
			} else {
				sorted = append(sorted, rv)
				j += 1
			}
		}
	}
	return sorted
}
