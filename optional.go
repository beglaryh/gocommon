package gocommon

import "errors"

type Optional[T any] struct {
	value *T
}

func WithPointer[T any](t *T) Optional[T] {
	return Optional[T]{value: t}
}

func With[T any](t T) Optional[T] {
	return Optional[T]{value: &t}
}

func Empty[T any]() Optional[T] {
	return Optional[T]{}
}

func (op Optional[T]) IsPresent() bool {
	return op.value != nil
}

func (op Optional[T]) IsEmpty() bool {
	return op.value == nil
}

func (op Optional[T]) GetPointer() (*T, error) {
	if op.value == nil {
		return op.value, errors.New("empty gocommon")
	}
	return op.value, nil
}

func (op Optional[T]) Get() (T, error) {
	if op.value == nil {
		var t T
		return t, errors.New("empty gocommon")
	}
	return *op.value, nil
}

func (op Optional[T]) OrElse(t T) *T {
	if !op.IsPresent() {
		return &t
	}
	return op.value
}

func (op Optional[T]) OrElseValue(t T) T {
	if !op.IsPresent() {
		return t
	}
	return *op.value
}

func (op Optional[T]) IfPresent(doSomething func(t T)) {
	if op.IsPresent() {
		doSomething(*op.value)
	}
}

func (op Optional[T]) IfPresentOrElse(doSomething func(t T), orElse func()) {
	if op.IsPresent() {
		doSomething(*op.value)
	} else {
		orElse()
	}
}
