package stream

import (
	"reflect"
	"testing"
)

type person struct {
	name string
	age  int
}

func TestStream_Filter(t *testing.T) {
	numbers := []int{1, 2, 3}
	filteredNumbers := Of[int](numbers).
		Filter(func(n int) bool { return n > 1 }).
		Slice()
	if 2 != filteredNumbers[0] {
		t.Fail()
	}
	if 3 != filteredNumbers[1] {
		t.Fail()
	}

	if 2 != len(filteredNumbers) {
		t.Fail()
	}
}

func TestStructFilter(t *testing.T) {
	//persons := []person{{"Bob", 50}, {"Bob", 100}, {"Alice", 30}}
	//filteredNumbers := gocommon.Of[person](persons).Filter(func(p person) bool { return p.name == "Bob" }).Slice()
	//assert.Equal(t, person{"Bob", 50}, filteredNumbers[0])
	//assert.Equal(t, person{"Bob", 100}, filteredNumbers[1])
	//assert.Equal(t, 2, len(filteredNumbers))

	//filteredNumbers = gocommon.Of[person](persons).Filter(func(p person) bool {
	//	return p.name == "Bob" && p.age == 100
	//}).Slice()
	//assert.Equal(t, person{"Bob", 100}, filteredNumbers[0])
	//assert.Equal(t, 1, len(filteredNumbers))
}

func TestStream_Map(t *testing.T) {
	persons := []person{{"Bob", 50}, {"Bob", 100}, {"Alice", 30}}
	names := Map[person, string](persons, func(p person) string {
		return p.name
	}).Filter(func(name string) bool { return name != "Bob" }).
		Slice()

	if "Alice" != names[0] {
		t.Fail()
	}
}

func TestStream_Sort(t *testing.T) {
	numbers := []int{5, 4, 3, 2, 1}
	sorted := Of(numbers).Sort(func(a, b int) bool { return a < b }).Slice()
	expected := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(expected, sorted) {
		t.Fail()
	}
	//assert.Equal(t, , sorted)

	persons := []person{{"Bob", 100}, {"Bob", 50}, {"Alice", 30}}
	sortedPersons := Of[person](persons).Sort(func(a, b person) bool {
		if a.name == b.name {
			return a.age < b.age
		}
		return a.name < b.name
	}).Slice()

	expected2 := []person{{"Alice", 30}, {"Bob", 50}, {"Bob", 100}}
	if !reflect.DeepEqual(expected2, sortedPersons) {
		t.Fail()
	}
}

func TestStream_AnyMatch(t *testing.T) {
	//numbers := []int{5, 4, 3, 2, 1}
	//match := Of(numbers).AnyMatch(func(a int) bool { return a == 1 })
	//assert.Equal(t, true, match)
	//
	//match = Of(numbers).AnyMatch(func(a int) bool { return a == 6 })
	//assert.Equal(t, false, match)
}

func TestStream_NoneMatch(t *testing.T) {
	//numbers := []int{5, 4, 3, 2, 1}
	//match := Of(numbers).NoneMatch(func(a int) bool { return a == 1 })
	//assert.Equal(t, false, match)
	//
	//match = Of(numbers).NoneMatch(func(a int) bool { return a == 6 })
	//assert.Equal(t, true, match)
}

func TestStream_FindFirst(t *testing.T) {
	numbers := []int{5, 4, 3, 2, 1}
	first := Of(numbers).FindFirst()
	v, _ := first.Get()
	if 5 != v {
		t.Fail()
	}

	first = Of(numbers).Filter(func(i int) bool { return i > 5 }).FindFirst()
	if first.IsPresent() {
		t.Fail()
	}
}

func TestFlatMap(t *testing.T) {
	numbers := [][]int{{1, 2, 3}, {4, 5, 6}}
	flattened := FlatMap[int](numbers).Slice()
	expected := []int{1, 2, 3, 4, 5, 6}
	if !reflect.DeepEqual(expected, flattened) {
		t.Fail()
	}
}

func TestGroupBy(t *testing.T) {
	persons := []person{{"Bob", 100}, {"Bob", 50}, {"Alice", 30}}
	group := GroupBy[string, person](persons, func(p person) string { return p.name })
	if 2 != len(group["Bob"]) {
		t.Fail()
	}
	if 1 != len(group["Alice"]) {
		t.Fail()
	}
}

func TestStream_ForEach(t *testing.T) {
	forEach := []string{}
	Of[string]([]string{"Hello", "World"}).ForEach(func(s string) {
		forEach = append(forEach, s)
	})
	expected := []string{"Hello", "World"}
	if !reflect.DeepEqual(expected, forEach) {
		t.Fail()
	}
}

func TestStream_Reduce(t *testing.T) {
	numbers := []int{1, 2, 3}
	sum := Of[int](numbers).Reduce(0, func(a, b int) int { return a + b })
	if 6 != sum {
		t.Fail()
	}
}
