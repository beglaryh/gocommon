package gocommon

import (
	"testing"
)

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

func TestOffsetDateTime(t *testing.T) {
	utc := "2024-04-03T06:14:36Z"
	ot, err := ParseOffsetDateTime(utc)
	if err != nil || ot.String() != utc {
		t.Fail()
	}

	nutc := "2007-12-03T10:15:30+01:00"
	ot, err = ParseOffsetDateTime(nutc)
	if err != nil || ot.String() != nutc {
		t.Fail()
	}
}

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
