package heap

import (
	"fmt"
	"testing"
)

func TestIndex(t *testing.T) {
	var index int

	index = 0
	expect(left(index), 1, t)
	expect(right(index), 2, t)
	expect(parent(index), 0, t)

	index = 1
	expect(left(index), 3, t)
	expect(right(index), 4, t)
	expect(parent(index), 0, t)

	index = 2
	expect(left(index), 5, t)
	expect(right(index), 6, t)
	expect(parent(index), 0, t)

	index = 4
	expect(left(index), 9, t)
	expect(right(index), 10, t)
	expect(parent(index), 1, t)
}

func expect(a, b int, t *testing.T) {
	if a != b {
		t.Fail()
	}
}

func TestMaxHeapify(t *testing.T) {
	h := newHeap([]interface{}{1, 2, 3, 4}, func(lower, topper interface{}) bool {
		return topper.(int) >= lower.(int)
	})
	h.Print()
	h.heapify(0)
	h.Print()
	if h.ValidCheck() == false {
		t.Fail()
	}
}

func TestMinHeapify(t *testing.T) {
	h := newHeap([]interface{}{7, 6, 5, 4}, func(lower, topper interface{}) bool {
		return topper.(int) <= lower.(int)
	})
	h.Print()
	h.heapify(0)
	h.Print()
	if h.ValidCheck() == false {
		t.Fail()
	}
}

func TestResetValueMaxHeap(t *testing.T) {
	h := NewHeap([]interface{}{1, 2, 3, 4, 5, 4, 3, 2, 1}, func(lower, topper interface{}) bool {
		return topper.(int) >= lower.(int)
	})
	h.Print()
	h.ResetIndex(2, 100)
	h.Print()
	if h.ValidCheck() == false {
		t.Fail()
	}
	h.Print()
	h.ResetIndex(0, 1)
	h.Print()
	if h.ValidCheck() == false {
		t.Fail()
	}
}

func TestResetValueMinHeap(t *testing.T) {
	h := NewHeap([]interface{}{1, 20, 3, 4, 5, 4, 3, 2, 1}, func(lower, topper interface{}) bool {
		return topper.(int) <= lower.(int)
	})
	h.Print()
	h.ResetIndex(2, 2)
	h.Print()
	if h.ValidCheck() == false {
		t.Fail()
	}
	h.Print()
	h.ResetIndex(0, 10)
	h.Print()
	if h.ValidCheck() == false {
		t.Fail()
	}
}

func TestPop(t *testing.T) {
	h := NewHeap([]interface{}{1, 2, 3, 4, 5, 4, 3, 2, 1}, func(lower, topper interface{}) bool {
		return lower.(int) >= topper.(int)
	})
	h.Print()
	top := h.PopTop()
	fmt.Printf("pop top:%d\n", top)
	h.Print()
	if h.ValidCheck() == false {
		t.Fail()
	}
}
