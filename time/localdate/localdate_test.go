package localdate

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func TestLocalDate(t *testing.T) {
	date := New(2024, 1, 1)
	if "2024-01-01" != date.String() {
		t.Fail()
	}

	date = date.PlusDays(1)
	if "2024-01-02" != date.String() {
		t.Fail()
	}

	date = New(2024, 5, 31)
	date = date.PlusMonths(1)
	if "2024-06-30" != date.String() {
		t.Fatalf("Expected: 2024-06-30. Actual: %s", date.String())
	}

	date = New(2024, 7, 31)
	date = date.MinusMonths(1)
	if "2024-06-30" != date.String() {
		t.Fatalf("Expected: 2024-06-30. Actual: %s", date.String())
	}
}

func TestBeforeAfter(t *testing.T) {
	date := Now()
	tomorrow := date.PlusDays(1)
	if !date.IsBefore(tomorrow) {
		t.Fail()
	}
	if !tomorrow.IsAfter(date) {
		t.Fail()
	}
}

func TestMarshalling(t *testing.T) {
	m := map[string]LocalDate{"v1": New(2024, time.March, 31)}
	js, _ := json.Marshal(m)
	var nm map[string]LocalDate
	json.Unmarshal(js, &nm)

	if !reflect.DeepEqual(m, nm) {
		t.Fail()
	}
}
