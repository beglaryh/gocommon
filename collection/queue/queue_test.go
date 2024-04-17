package queue

import "testing"

func TestFifoQueue(t *testing.T) {
	fifo := NewFifoQueue[int]()
	fifo.Add(1, 2, 3, 4)

	if fifo.Size() != 4 {
		t.Fail()
	}

	peek, err := fifo.Peek().Get()

	if err != nil || peek != 1 {
		t.Fail()
	}

	pop, err := fifo.Remove().Get()
	if err != nil || pop != 1 {
		t.Fail()
	}

}
