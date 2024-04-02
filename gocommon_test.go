package gocommon

import "testing"

func TestList(t *testing.T) {
	l := EmptyList[int]()
	if l.Size() != 0 {
		t.Fail()
	}
	l.Add(1)
	if l.Size() != 1 {
		t.Fail()
	}
	if l.Get(0).Get() != 1 {
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

}
