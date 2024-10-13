package priorityqueue

import (
	"testing"

	"github.com/beglaryh/gocommon/collection/list/arraylist"
)

func TestMaxHeap(t *testing.T) {
	queue := New[int](func(a, b int) int {
		return a - b
	})

	queue.Add(1)
	peek, _ := queue.Peek()
	if peek != 1 {
		t.Fail()
	}
	queue.Add(2)
	peek, _ = queue.Peek()
	if peek != 2 {
		t.Fail()
	}
	queue.Add(3)
	peek, _ = queue.Peek()
	if peek != 3 {
		t.Fail()
	}
	queue.Add(4)
	peek, _ = queue.Peek()
	if peek != 4 {
		t.Fail()
	}
}

func TestPollEmpty(t *testing.T) {
	pq := New[int](maxHeap)
	_, err := pq.Poll()
	if err == nil {
		t.Fail()
	}
}

func TestPollSingle(t *testing.T) {
	pq := New[int](maxHeap)
	pq.Add(1)
	e, _ := pq.Poll()
	if e != 1 || !pq.IsEmpty() {
		t.Fail()
	}
}

func TestPollTwice(t *testing.T) {
	pq := New[int](maxHeap)
	pq.Add(1, 2)
	e, _ := pq.Poll()
	if e != 2 {
		t.Fail()
	}
	e, _ = pq.Poll()
	if e != 1 {
		t.Fail()
	}

	_, err := pq.Poll()
	if err == nil {
		t.Fail()
	}
}

func TestPollFour(t *testing.T) {
	pq := New[int](maxHeap)
	pq.Add(1, 2, 3, 4)
	expected := arraylist.New[int]()
	expected.Add(4, 3, 2, 1)
	result := arraylist.New[int]()
	for !pq.IsEmpty() {
		v, _ := pq.Poll()
		result.Add(v)
	}

	if !expected.Equals(result) {
		t.Fail()
	}
}

func TestPollMany(t *testing.T) {
	pq := New[int](maxHeap)
	expected := arraylist.New[int]()
	for i := range 100 {
		pq.Add(i)
		expected.Add(100 - i - 1)
	}

	result := arraylist.New[int]()
	for !pq.IsEmpty() {
		v, _ := pq.Poll()
		result.Add(v)
	}

	if !expected.Equals(result) {
		t.Fail()
	}
}

func maxHeap(a, b int) int {
	return a - b
}
