package heap

import (
	"fmt"
	"testing"
)

func TestIndex(t *testing.T) {
	var index int

	index = 0
	expect(leftIndex(index), 1, t)
	expect(rightIndex(index), 2, t)
	expect(parentIndex(index), 0, t)

	index = 1
	expect(leftIndex(index), 3, t)
	expect(rightIndex(index), 4, t)
	expect(parentIndex(index), 0, t)

	index = 2
	expect(leftIndex(index), 5, t)
	expect(rightIndex(index), 6, t)
	expect(parentIndex(index), 0, t)

	index = 4
	expect(leftIndex(index), 9, t)
	expect(rightIndex(index), 10, t)
	expect(parentIndex(index), 1, t)
}

func expect(a, b int, t *testing.T) {
	if a != b {
		t.Fail()
	}
}

func TestMaxHeapify(t *testing.T) {
	h := NewHeap([]interface{}{1, 2, 3, 4}, func(lower, topper interface{}) bool {
		return topper.(int) >= lower.(int)
	})
	h.print()
	h.heapify(0)
	h.print()
	h.buildHeap()
	h.print()
	if h.check() == false {
		t.Fail()
	}
}

func TestMinHeapify(t *testing.T) {
	h := NewHeap([]interface{}{7, 6, 5, 4}, func(lower, topper interface{}) bool {
		return topper.(int) >= lower.(int)
	})
	h.print()
	h.heapify(0)
	h.print()
	h.buildHeap()
	h.print()
	if h.check() == false {
		t.Fail()
	}
}

func TestResetValueMaxHeap(t *testing.T) {
	h := NewHeap([]interface{}{1, 2, 3, 4, 5, 4, 3, 2, 1}, func(lower, topper interface{}) bool {
		return topper.(int) >= lower.(int)
	})
	h.print()
	h.ResetIndex(2, 100)
	h.print()
	if h.check() == false {
		t.Fail()
	}
	h.print()
	h.ResetIndex(0, 1)
	h.print()
	if h.check() == false {
		t.Fail()
	}
}

func TestResetValueMinHeap(t *testing.T) {
	h := NewHeap([]interface{}{1, 20, 3, 4, 5, 4, 3, 2, 1}, func(lower, topper interface{}) bool {
		return topper.(int) >= lower.(int)
	})
	h.buildHeap()
	h.print()
	h.ResetIndex(2, 2)
	h.print()
	if h.check() == false {
		t.Fail()
	}
	h.print()
	h.ResetIndex(0, 10)
	h.print()
	if h.check() == false {
		t.Fail()
	}
}

func TestPop(t *testing.T) {
	h := NewHeap([]interface{}{1, 2, 3, 4, 5, 4, 3, 2, 1}, func(lower, topper interface{}) bool {
		return lower.(int) >= topper.(int)
	})
	h.print()
	top := h.popTop()
	fmt.Printf("pop top:%d\n", top)
	h.print()
	if h.check() == false {
		t.Fail()
	}
}
