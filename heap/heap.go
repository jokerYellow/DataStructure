package heap

import (
	"fmt"
	"math"
)

type HeapType int

const (
	minHeap = HeapType(1)
	maxHeap = HeapType(2)
)

type Heap struct {
	heapType                    HeapType
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

func (h *Heap) ResetIndex(index int, value interface{}) {
	if index >= h.heapSize {
		return
	}
	origin := h.items[index]
	h.items[index] = value

	originIsBigger := h.whetherIndexIsBiggerOrEqual(value, origin)
	switch h.heapType {
	case minHeap:
		if originIsBigger {
			h.up(index)
		} else {
			h.down(index)
		}
		break
	case maxHeap:
		if originIsBigger {
			h.down(index)
		} else {
			h.up(index)
		}
		break
	}
}

func (h *Heap) up(index int) {
	if index > h.heapSize {
		return
	}
	for index > 0 {
		switch h.heapType {
		case minHeap:
			if h.whetherIndexIsBiggerOrEqual(h.items[index], h.items[parentIndex(index)]) == false {
				return
			}
			break
		case maxHeap:
			if h.whetherIndexIsBiggerOrEqual(h.items[index], h.items[parentIndex(index)]) {
				return
			}
			break
		}
		h.swap(parentIndex(index), index)
		index = parentIndex(index)
	}
}

func (h *Heap) swap(i, j int) {
	if i >= h.heapSize || j >= h.heapSize {
		return
	}
	h.items[j], h.items[i] = h.items[i], h.items[j]
}

func (h *Heap) down(index int) {
	if index > h.heapSize {
		return
	}
	h.heapify(index)
	return
	for {
		switch h.heapType {
		case minHeap:
			if h.whetherIndexIsBiggerOrEqual(h.items[index], h.items[parentIndex(index)]) == false {
				return
			}
			lIndex := leftIndex(index)
			rIndex := rightIndex(index)
			if lIndex > h.heapSize {
				return
			}
			largestIndex := index
			if h.whetherIndexIsBiggerOrEqual(h.items[index], h.items[lIndex]) {
				largestIndex = lIndex
			}
			if rIndex <= h.heapSize {
				if h.whetherIndexIsBiggerOrEqual(h.items[largestIndex], h.items[rIndex]) {
					largestIndex = rIndex
				}
			}
			h.swap(largestIndex, index)
			index = largestIndex
			break
		case maxHeap:
			if h.whetherIndexIsBiggerOrEqual(h.items[index], h.items[parentIndex(index)]) {
				return
			}
			lIndex := leftIndex(index)
			rIndex := rightIndex(index)
			if lIndex > h.heapSize {
				return
			}
			maxIndex := index
			if h.whetherIndexIsBiggerOrEqual(h.items[index], h.items[lIndex]) {
				maxIndex = lIndex
			}
			if rIndex <= h.heapSize {
				if h.whetherIndexIsBiggerOrEqual(h.items[maxIndex], h.items[rIndex]) {
					maxIndex = rIndex
				}
			}
			h.swap(maxIndex, index)
			index = maxIndex
			break
		}
	}
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
	expectResult := false
	if h.heapType == minHeap {
		expectResult = true
	}
	largestIndex := index
	if h.whetherIndexIsBiggerOrEqual(h.items[lIndex], cItem) == expectResult {
		largestIndex = lIndex
	}
	if rIndex < h.heapSize &&
		h.whetherIndexIsBiggerOrEqual(h.items[rIndex], h.items[largestIndex]) == expectResult {
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
		expectValue := false
		//TODO:expectValue is wrong
		if lIndex < h.heapSize {
			if h.whetherIndexIsBiggerOrEqual(h.items[lIndex], v) == expectValue {
				return false
			}
		}
		if rIndex < h.heapSize {
			if h.whetherIndexIsBiggerOrEqual(h.items[rIndex], v) == expectValue {
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
