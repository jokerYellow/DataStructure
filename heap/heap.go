package heap

import (
	"fmt"
	"math"
)

type heap struct {
	items []interface{}
	/*left is bigger than right*/
	whetherIndexIsBiggerOrEqual func(other, index interface{}) bool
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

func (h *heap) heapSize() int {
	return len(h.items)
}

func (h *heap) heapify(index int) {
	if index >= h.heapSize() {
		return
	}
	cItem := h.items[index]
	lIndex := leftIndex(index)
	rIndex := rightIndex(index)
	if lIndex < h.heapSize() {
		lItem := h.items[lIndex]
		largestIndex := index
		if h.whetherIndexIsBiggerOrEqual(lItem, cItem) == false {
			largestIndex = lIndex
		}
		if rIndex < h.heapSize() {
			rItem := h.items[rIndex]
			if h.whetherIndexIsBiggerOrEqual(rItem, h.items[largestIndex]) == false {
				largestIndex = rIndex
			}
		}
		if largestIndex != index {
			h.items[index], h.items[largestIndex] = h.items[largestIndex], h.items[index]
			h.heapify(largestIndex)
		}
	}
}

func (h *heap) buildHeap() {
	i := h.heapSize() / 2
	for index := i; index >= 0; index-- {
		h.heapify(index)
	}
}

func (h *heap) check() bool {
	for i, v := range h.items {
		lIndex := leftIndex(i)
		rIndex := rightIndex(i)
		if lIndex < h.heapSize() {
			if h.whetherIndexIsBiggerOrEqual(h.items[lIndex], v) == false {
				return false
			}
		}
		if rIndex < h.heapSize() {
			if h.whetherIndexIsBiggerOrEqual(h.items[rIndex], v) == false {
				return false
			}
		}
	}
	return true
}

func (h *heap) print() {
	bounds := 2

	rows := int(math.Ceil(math.Log2(float64(h.heapSize())))) - 1
	rows = int(math.Pow(2, float64(rows)))

	for index, value := range h.items {
		for i := 0; i < int(rows); i++ {
			fmt.Printf(" ")
		}
		fmt.Printf("%d", value)
		if index == bounds-2 {
			fmt.Printf("\n")
			bounds = bounds * 2
			rows = rows / 2
		}
	}
	fmt.Printf("\n")
}
