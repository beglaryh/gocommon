package gocommon

import (
	"github.com/beglaryh/gocommon/optional"
	"testing"
)

func TestOptional(t *testing.T) {
	a := 1
	op := optional.With[int](a)
	if op.IsEmpty() {
		t.Fail()
	}

	b, _ := op.GetPointer()
	if *b != a {
		t.Fail()
	}

	op = optional.WithPointer[int](nil)
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
