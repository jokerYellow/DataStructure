package heap

import (
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
	h := &heap{}
	h.items = []interface{}{1, 2, 3, 4, 5, 6, 7}
	//this is a maxHeap
	h.whetherIndexIsBiggerOrEqual = func(other, index interface{}) bool {
		itemValue := other.(int)
		indexValue := index.(int)
		return indexValue >= itemValue
	}
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
	h := &heap{}
	h.items = []interface{}{7, 6, 5, 4, 3, 2, 1}
	//this is a maxHeap
	h.whetherIndexIsBiggerOrEqual = func(other, index interface{}) bool {
		itemValue := other.(int)
		indexValue := index.(int)
		return indexValue <= itemValue
	}
	h.print()
	h.heapify(0)
	h.print()
	h.buildHeap()
	h.print()
	if h.check() == false {
		t.Fail()
	}
}