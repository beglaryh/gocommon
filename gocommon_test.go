package gocommon

import (
	"testing"
)

func TestOptional(t *testing.T) {
	a := 1
	op := With[int](a)
	if op.IsEmpty() {
		t.Fail()
	}

	b, _ := op.GetPointer()
	if *b != a {
		t.Fail()
	}

	op = WithPointer[int](nil)
	if op.IsPresent() {
		t.Fail()
	}

	_, err := op.Get()
	if err == nil {
		t.Fail()
	}

	_, err = op.GetPointer()
	if err == nil {
		t.Fail()
	}
}
