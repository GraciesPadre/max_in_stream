package main

import (
	"strconv"
	"strings"
)

/*
Assumptions:
 - "Given an infinite stream of numbers, make a function to find the top 50 integers."  I took this to mean that
   what we want are the top 50 unique values.  For example, given a stream of 1, 1, 1, we would want a collection
   with one entry of value 1.
 - The resulting collection need not be sorted, though I did sort the resulting collection for use in testing.
 - I interpreted infinite to mean 2 million, just so that the program would eventually finish and yield a result.

I considered 2 options when implementing:
 - Use an array that is always sorted in ascending order.  When adding a new value, if that new value is <=
   array[0], skip it.  Otherwise, push that value to array[0], sort, and, if there are more elements in the
   array than we want to keep, pop array[0].  Was simple and worked, but seemed like a lot of sorting.

What I decided to use is a min heap implemented in terms of an array.  Since a min heap will always have the root
node contain the smallest value, and that value is in the collection in array[0], you can skip a new value if it
is <= array[0].  Doing so prevents having duplicate max values.  I used the go runtime library min heap
(https://cs.opensource.google/go/go/+/refs/tags/go1.17.1:src/container/heap/heap.go)
as the basis for the min heap implementation and added:
 - Skipping a new value by comparing it to array[0] to avoid duplicated max values.
 - Popping the heap if adding a new value causes the collection size to exceed the number of values we want to
   store.

With the go 1.15 development tools installed:
 - Run from the folder containing main.go using the shell command
   'go run max_in_stream'
   Doing so should yield a reult like this...
   999951,999952,999959,999968,999953,999961,999960,999970,999974,999955,999954,999963,999962,999965,999966,999982,999971,999976,999975,999956,999967,999957,999978,999990,999964,999985,999987,999984,999983,999999,999996,999995,999988,999989,999994,999986,999992,999993,999977,1000000,999972,999969,999973,999958,999980,999979,999981,999997,999998,999991
   The value 999951 should always be first value printed.  The values between 999952 and 1000000 should be in the
   result but in no particular order.
 - Test using the shell command
   'go test'
 - There is a prebuilt MacOs binary named max_in_stream you can run from a shell in the github repo.
 */

type MinimumHeap struct {
	numbers []int
	numValuesToKeep int
}

func NewMinimumHeap(numValuesToKeep int) *MinimumHeap {
	return &MinimumHeap{
		numbers: []int{},
		numValuesToKeep: numValuesToKeep,
	}
}

func (heap *MinimumHeap) Count() int {
	return len(heap.numbers)
}

func (heap *MinimumHeap) Push(value int) {
	if len(heap.numbers) > 0 && value <= heap.numbers[0] {
		return
	}

	heap.numbers = append(heap.numbers, value)
	heap.up(len(heap.numbers) - 1)

	if heap.Count() > heap.numValuesToKeep {
		heap.Pop()
	}
}

func (heap *MinimumHeap) up(index int) {
	for {
		parentIdx := (index - 1) / 2
		if index == 0 || heap.numbers[parentIdx] <= heap.numbers[index] {
			break
		}
		heap.numbers[index], heap.numbers[parentIdx] = heap.numbers[parentIdx], heap.numbers[index]
		index = parentIdx
	}
}

func (heap *MinimumHeap) Pop() (int, bool) {
	if len(heap.numbers) == 0 {
		return 0, false
	}

	value := heap.numbers[0]
	heap.numbers[0], heap.numbers[len(heap.numbers)-1] = heap.numbers[len(heap.numbers)-1], heap.numbers[0]
	heap.numbers = heap.numbers[:len(heap.numbers)-1]
	heap.down(0)

	return value, true
}

func (heap *MinimumHeap) down(index int) {
	for {
		// pick child to swap (smaller one)
		childIndex := index*2 + 1                              // left child
		if childIndex >= len(heap.numbers) || childIndex < 0 { // <0 int overflow
			break
		}
		rightIndex := childIndex + 1
		if rightIndex < len(heap.numbers) && heap.numbers[childIndex] >= heap.numbers[rightIndex] {
			childIndex = rightIndex
		}

		// swap
		if heap.numbers[childIndex] >= heap.numbers[index] {
			break
		}
		heap.numbers[index], heap.numbers[childIndex] = heap.numbers[childIndex], heap.numbers[index]
		index = childIndex
	}
}

func (heap *MinimumHeap) String() string {
	var text []string

	for _, number := range heap.numbers {
		text = append(text, strconv.Itoa(number))
	}

	return strings.Join(text, ",")
}
