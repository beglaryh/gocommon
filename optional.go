package optional

import "errors"

type Optional[T any] struct {
	value *T
}

func With[T any](t *T) (*Optional[T], error) {
	if t == nil {
		return nil, errors.New("nil argument passed")
	}
	return &Optional[T]{value: t}, nil
}

func WithNullable[T any](t *T) *Optional[T] {
	return &Optional[T]{value: t}
}

func Empty[T any]() Optional[T] {
	return Optional[T]{}
}

func (op *Optional[T]) IsPresent() bool {
	return op.value != nil
}

func (op *Optional[T]) Get() T {
	return *op.value
}

func (op *Optional[T]) IfPresent(doSomething func(t T)) {
	if op.IsPresent() {
		doSomething(*op.value)
	}
}

func (op *Optional[T]) IfPresentOrElse(doSomething func(t T), orElse func()) {
	if op.IsPresent() {
		doSomething(*op.value)
	} else {
		orElse()
	}
}
