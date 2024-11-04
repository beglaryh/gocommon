package arraylist

import (
	"errors"
	"reflect"
	"testing"

	"github.com/beglaryh/gocommon/collection/collection_errors"
)

func TestArrayList(t *testing.T) {
	l := New[int]()
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

	l = New[int]()
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
	_, err := NewBuilder[int]().WithElements([]int{1, 2}).WithLimit(1).Build()
	if !errors.Is(err, collection_errors.LimitExceeded) {
		t.Fail()
	}

	l, _ := NewBuilder[int]().WithLimit(1).Build()
	_ = l.Add(1)
	err = l.Add(2)
	if !errors.Is(err, collection_errors.LimitExceeded) {
		t.Fail()
	}

	l, _ = NewBuilder[int]().WithElements([]int{1, 2, 3, 4}).Build()
	if !reflect.DeepEqual([]int{1, 2, 3, 4}, l.ToArray()) {
		t.Fail()
	}
}

func TestIter(t *testing.T) {
	al := New[int]()
	al.Add(10)
	al.Add(20)
	al.Add(30)

	iSumExpected := 3
	vSumExpected := 60
	var iSum int
	var vSum int
	for i, v := range al.Iter {
		iSum += i
		vSum += v
	}

	if vSum != vSumExpected || iSum != iSumExpected {
		t.Fail()
	}

}
