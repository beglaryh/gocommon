package priorityqueue

import (
	"testing"
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
