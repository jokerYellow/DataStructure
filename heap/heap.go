package heap

import (
	"fmt"
	"math"
)

type Heaper interface {
	Top() interface{}
	ResetIndex(index int, value interface{})
	PopTop() interface{}
	ValidCheck() bool
	Print()
}

type Heap struct {
	items    []interface{}
	heapSize int
	compare  Compare
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func parent(i int) int {
	return (i - 1) / 2
}

type Compare func(lower, topper interface{}) bool

func NewHeap(items []interface{}, compare Compare) Heaper {
	return newHeap(items, compare)
}

func newHeap(items []interface{}, compare Compare) *Heap {
	h := new(Heap)
	h.items = items
	h.compare = compare
	h.heapSize = len(h.items)
	h.buildHeap()
	return h
}

func (h *Heap) Top() interface{} {
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

	//judge is minHeap or maxHeap
	//if heapSize is 0 or 1,either min or max is ok.

	flag := true
	if h.heapSize >= 2 {
		flag = h.compare(h.items[1], h.items[0])
		return
	}
	h.items[index] = value
	if h.compare(value, origin) != flag {
		h.up(index)
	} else {
		h.down(index)
	}
}

func (h *Heap) up(index int) {
	if index > h.heapSize {
		return
	}
	for index > 0 {
		if h.compare(h.items[index], h.items[parent(index)]) {
			return
		}
		h.swap(parent(index), index)
		index = parent(index)
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
}

func (h *Heap) PopTop() interface{} {
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
	lIndex := left(index)
	rIndex := right(index)
	if lIndex >= h.heapSize {
		return
	}
	largestIndex := index
	if h.compare(h.items[lIndex], cItem) == false {
		largestIndex = lIndex
	}
	if rIndex < h.heapSize &&
		h.compare(h.items[rIndex], h.items[largestIndex]) == false {
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

func (h *Heap) ValidCheck() bool {
	for i, v := range h.items {
		lIndex := left(i)
		rIndex := right(i)
		if lIndex < h.heapSize {
			if h.compare(h.items[lIndex], v) == false {
				return false
			}
		}
		if rIndex < h.heapSize {
			if h.compare(h.items[rIndex], v) == false {
				return false
			}
		}
	}
	return true
}

func (h *Heap) Print() {
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
