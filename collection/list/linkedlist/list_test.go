package linkedlist

import (
	"testing"
)

func TestLinkedList(t *testing.T) {

	l := New[string]()
	l.Add("A")
	if l.Size() != 1 {
		t.Fail()
	}
	v, _ := l.Get(0)
	if v != "A" {
		t.Fail()
	}

	v, _ = l.Get(-1)
	if v != "A" {
		t.Fail()
	}

	l.Add("B")

	if l.Size() != 2 {
		t.Fail()
	}

	v, _ = l.Get(1)
	if v != "B" {
		t.Fail()
	}

	v, _ = l.Get(-1)
	if v != "B" {
		t.Fail()
	}

}
