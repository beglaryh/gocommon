package arraylist

import (
	"errors"
	"github.com/beglaryh/gocommon/collection/collection_errors"
	"reflect"
	"testing"
)

func TestArrayList(t *testing.T) {
	l, _ := New[int]()
	if l.Size() != 0 {
		t.Fail()
	}
	l.Add(1)
	if l.Size() != 1 {
		t.Fail()
	}

	e, _ := l.Get(0)
	if e != 1 {
		t.Fail()
	}

	l.Clear()
	if l.Size() != 0 {
		t.Fail()
	}

	for i := range 100 {
		l.Add(i)
	}

	if l.Size() != 100 {
		t.Fail()
	}

	if 100 != len(l.ToArray()) {
		t.Fail()
	}

	l, _ = New[int]()
	_ = l.Add(1)
	_ = l.Add(2)
	e, _ = l.Remove(0)
	if e != 1 && l.Size() != 1 {
		t.Fail()
	}

	e, _ = l.Remove(-1)
	if e != 2 && l.IsEmpty() {
		t.Fail()
	}

	_, err := l.Remove(0)
	if err == nil {
		t.Fail()
	}
}

func TestBuilder(t *testing.T) {
	l, err := NewBuilder[int]().WithElements([]int{1, 2}).WithLimit(1).Build()
	if !errors.Is(err, collection_errors.LimitExceeded) {
		t.Fail()
	}

	l, _ = NewBuilder[int]().WithLimit(1).Build()
	_ = l.Add(1)
	err = l.Add(2)
	if !errors.Is(err, collection_errors.LimitExceeded) {
		t.Fail()
	}

	l, _ = NewBuilder[int]().WithElements([]int{1, 2, 3, 4}).Build()
	if !reflect.DeepEqual([]int{1, 2, 3, 4}, l.elements) {
		t.Fail()
	}
}
