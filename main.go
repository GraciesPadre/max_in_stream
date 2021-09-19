package main

import "fmt"

func main() {
	const numValuesToKeep = 50
	const start = -999_999
	const end = 1_000_000

	heap := NewMinimumHeap(numValuesToKeep)

	for i := start; i <= end; i++ {
		heap.Push(i)
	}

	fmt.Println(heap)
}
