package main

import "github.com/jokerYellow/DataStructure/heap"

func main() {
	h := heap.NewHeap([]interface{}{100, 1, 1, 1, 1, 1, 1, 1, 2, 3}, func(lower, topper interface{}) bool {
		return lower.(int) <= topper.(int)
	})
	h.Print()
}
