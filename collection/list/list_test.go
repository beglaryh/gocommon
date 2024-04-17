package list

import (
	. "github.com/beglaryh/gocommon"
	"testing"
)

func TestArrayList(t *testing.T) {
	l, _ := NewArrayList[int]()
	if l.Size() != 0 {
		t.Fail()
	}
	l.Add(1)
	if l.Size() != 1 {
		t.Fail()
	}

	e, _ := l.Get(0).Get()
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

	anotherList, err := NewArrayList(ArrayListWithSlice(l.ToArray()))
	if err != nil {
		t.Fail()
	}
	if 100 != anotherList.Size() {
		t.Fail()
	}

	l, _ = NewArrayList[int](ArrayListWithLimit[int](1))
	_ = l.Add(1)
	err = l.Add(2)
	if err == nil {
		t.Fail()
	}
}

func TestLinkedList(t *testing.T) {

	l := NewLinkedList[String]()
	a := String("A")
	l.Add(a)
	if l.Size() != 1 {
		t.Fail()
	}
	v, _ := l.Get(0).Get()
	if v != String("A") {
		t.Fail()
	}

	v, _ = l.Get(-1).Get()
	if v != String("A") {
		t.Fail()
	}

	l.Add(String("B"))

	if l.Size() != 2 {
		t.Fail()
	}

	v, _ = l.Get(1).Get()
	if v != String("B") {
		t.Fail()
	}

	v, _ = l.Get(-1).Get()
	if v != String("B") {
		t.Fail()
	}

}
