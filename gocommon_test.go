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

	anotherList := NewList(l.ToArray())
	if 100 != anotherList.Size() {
		t.Fail()
	}
}

func TestLocalDate(t *testing.T) {
	date := NewLocalDate(2024, 1, 1)
	if "2024-01-01" != date.String() {
		t.Fail()
	}

	date = date.PlusDays(1)
	if "2024-01-02" != date.String() {
		t.Fail()
	}

	date = NewLocalDate(2024, 5, 31)
	date = date.PlusMonths(1)
	if "2024-06-30" != date.String() {
		t.Fatalf("Expected: 2024-06-30. Actual: %s", date.String())
	}

	date = NewLocalDate(2024, 7, 31)
	date = date.MinusMonths(1)
	if "2024-06-30" != date.String() {
		t.Fatalf("Expected: 2024-06-30. Actual: %s", date.String())
	}
}
