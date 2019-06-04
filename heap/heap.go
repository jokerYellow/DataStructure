package heap

import (
	"fmt"
	"math"
)

type Heap struct {
	items                       []interface{}
	heapSize                    int
	whetherIndexIsBiggerOrEqual func(other, root interface{}) bool
}

func leftIndex(i int) int {
	return 2*i + 1
}

func rightIndex(i int) int {
	return 2*i + 2
}

func parentIndex(i int) int {
	return (i - 1) / 2
}

func NewHeap(items []interface{}, compare func(other, root interface{}) bool) Heap {
	h := Heap{}
	h.items = items
	h.heapSize = len(items)
	h.whetherIndexIsBiggerOrEqual = compare
	h.buildHeap()
	return h
}

func (h *Heap) top() interface{} {
	if len(h.items) == 0 {
		return nil
	}
	return h.items[0]
}

func (h *Heap) popTop() interface{} {
	if len(h.items) == 0 {
		return nil
	}
	rt := h.items[0]
	h.items[0] = h.items[h.heapSize-1]
	h.items[h.heapSize-1] = nil
	h.heapSize -= 1
	h.heapify(0)
	return rt
}

func (h *Heap) heapify(index int) {
	if index >= h.heapSize {
		return
	}
	cItem := h.items[index]
	lIndex := leftIndex(index)
	rIndex := rightIndex(index)
	if lIndex >= h.heapSize {
		return
	}
	largestIndex := index
	if h.whetherIndexIsBiggerOrEqual(h.items[lIndex], cItem) == false {
		largestIndex = lIndex
	}
	if rIndex < h.heapSize &&
		h.whetherIndexIsBiggerOrEqual(h.items[rIndex], h.items[largestIndex]) == false {
		largestIndex = rIndex
	}
	if largestIndex != index {
		h.items[index], h.items[largestIndex] = h.items[largestIndex], h.items[index]
		h.heapify(largestIndex)
	}
}

func (h *Heap) buildHeap() {
	i := h.heapSize / 2
	for index := i; index >= 0; index-- {
		h.heapify(index)
	}
}

func (h *Heap) check() bool {
	for i, v := range h.items {
		lIndex := leftIndex(i)
		rIndex := rightIndex(i)
		if lIndex < h.heapSize {
			if h.whetherIndexIsBiggerOrEqual(h.items[lIndex], v) == false {
				return false
			}
		}
		if rIndex < h.heapSize {
			if h.whetherIndexIsBiggerOrEqual(h.items[rIndex], v) == false {
				return false
			}
		}
	}
	return true
}

func (h *Heap) print() {
	bounds := 2

	rows := int(math.Ceil(math.Log2(float64(h.heapSize)))) - 1
	rows = int(math.Pow(2, float64(rows)))

	for index := 0; index < h.heapSize; index++ {
		for i := 0; i < int(rows); i++ {
			fmt.Printf(" ")
		}
		fmt.Printf("%d", h.items[index])
		if index == bounds-2 {
			fmt.Printf("\n")
			bounds = bounds * 2
			rows = rows / 2
		}
	}
	fmt.Printf("\n")
}
