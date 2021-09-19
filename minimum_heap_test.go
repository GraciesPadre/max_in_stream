package main

import (
	"sort"
	"testing"
)

func TestEmptyHeap(t *testing.T) {
	heap := NewMinimumHeap(50)
	if heap.Count() != 0 {
		t.Fatal("Count should have been 0")
	}
}

func TestHeapWithOneValue(t *testing.T) {
	heap := NewMinimumHeap(50)
	heap.Push(1)
	if heap.Count() != 1 {
		t.Fatal("Count should have been 1")
	}
}

func TestHeapWithSameValue(t *testing.T) {
	heap := NewMinimumHeap(50)

	heap.Push(1)
	heap.Push(2)
	heap.Push(1)

	if heap.Count() != 2 {
		t.Fatal("Count should have been 2")
	}

	sort.Ints(heap.numbers)
	if heap.numbers[0] != 1 {
		t.Fatal("numbers[0] should have been 1")
	}

	if heap.numbers[1] != 2 {
		t.Fatal("numbers[1] should have been 2")
	}
}

func TestALargeNumberOfValues(t *testing.T)  {
	const numValuesToKeep = 50
	const start = -999_999
	const end = 1_000_000

	heap := NewMinimumHeap(numValuesToKeep)

	for i := start; i <= end; i++ {
		heap.Push(i)
	}

	if heap.Count() != numValuesToKeep {
		t.Fatal("Count should have been 50")
	}

	sort.Ints(heap.numbers)

	if heap.numbers[0] != end - numValuesToKeep + 1 {
		t.Fatalf("heap.numbers[0] should have been %d, but was %d", end - numValuesToKeep + 1, heap.numbers[0])
	}

	if heap.numbers[numValuesToKeep - 1] != end {
		t.Fatalf("heap.numbers[%d] should have been %d, but was %d", numValuesToKeep - 1, end, heap.numbers[heap.numValuesToKeep -1])
	}
}

func TestALargeNumberOfValuesAlternatingSign(t *testing.T)  {
	const numValuesToKeep = 50
	const start = -999_999
	const end = 1_000_000

	heap := NewMinimumHeap(numValuesToKeep)


	for i := start; i <= end; i++ {
		if i % 2 == 0 {
			heap.Push(-i)
		} else {
			heap.Push(i)
		}
	}

	if heap.Count() != numValuesToKeep {
		t.Fatal("Count should have been 50")
	}

	sort.Ints(heap.numbers)

	if heap.numbers[0] != end - numValuesToKeep {
		t.Fatalf("heap.numbers[0] should have been %d, but was %d", end - numValuesToKeep, heap.numbers[0])
	}

	if heap.numbers[numValuesToKeep - 1] != end - 1{
		t.Fatalf("heap.numbers[%d] should have been %d, but was %d", numValuesToKeep - 1, end - 1, heap.numbers[heap.numValuesToKeep -1])
	}
}

func TestPopOfEmptyHeap(t *testing.T) {
	const numValuesToKeep = 50

	heap := NewMinimumHeap(numValuesToKeep)
	value, popped := heap.Pop()

	if value != 0 {
		t.Fatal("value should have been 0")
	}

	if popped {
		t.Fatal("popped should have been false")
	}
}

func TestPopOfHeapWith1Element(t *testing.T) {
	const numValuesToKeep = 50

	heap := NewMinimumHeap(numValuesToKeep)
	heap.Push(1)
	value, popped := heap.Pop()

	if value != 1 {
		t.Fatal("value should have been 1")
	}

	if !popped {
		t.Fatal("popped should have been true")
	}
}

func TestPopOfHeapWithManyElements(t *testing.T) {
	const numValuesToKeep = 50
	const start = -999_999
	const end = 1_000_000

	heap := NewMinimumHeap(numValuesToKeep)

	for i := start; i <= end; i++ {
		heap.Push(i)
	}

	value, popped := heap.Pop()

	if value != end - numValuesToKeep + 1 {
		t.Fatalf("value should have been %d", end - numValuesToKeep + 1)
	}

	if !popped {
		t.Fatal("popped should have been true")
	}

	if heap.Count() != numValuesToKeep -1 {
		t.Fatalf("Count should have been %d but was %d", numValuesToKeep - 1, heap.Count())
	}
}
