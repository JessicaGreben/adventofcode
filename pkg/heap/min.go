package heap

import (
	"errors"
	"fmt"
)

var (
	ErrEmptyHeap = errors.New("heap is empty")
)

type MinHeap[T any] struct {
	elements []T
	cmp      func(x, y T) bool
}

func NewMinHeap[T any](cmp func(x, y T) bool) *MinHeap[T] {
	return &MinHeap[T]{
		elements: []T{},
		cmp:      cmp,
	}
}

func (h MinHeap[T]) Print() {
	fmt.Println(h.elements)
}
func (h MinHeap[T]) Len() int {
	return len(h.elements)
}

func (h *MinHeap[T]) swap(i, j int) {
	h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
}

func (h *MinHeap[T]) Peak() T {
	if len(h.elements) < 1 {
		var t T
		return t
	}
	return h.elements[0]
}

func (h *MinHeap[T]) Push(x T) {
	h.elements = append(h.elements, x)
	h.bubbleUp()
}

// bubbleUp moves that last item in the heap up if its smaller than its parent.
func (h *MinHeap[T]) bubbleUp() {
	currIdx := h.Len() - 1
	parentIdx := (currIdx - 1) / 2
	for parentIdx >= 0 {
		p, c := h.elements[parentIdx], h.elements[currIdx]
		if h.cmp(c, p) {
			h.swap(parentIdx, currIdx)
			currIdx = parentIdx
			parentIdx = (currIdx - 1) / 2
		} else {
			return
		}
	}
}

func (h *MinHeap[T]) Pop() (T, error) {
	if h.Len() == 0 {
		var t T
		return t, ErrEmptyHeap
	}
	x := h.elements[0]
	h.swap(0, h.Len()-1)
	h.elements = h.elements[:h.Len()-1]
	h.bubbleDown()
	return x, nil
}

// bubbleDown moves the root of the heap down if its larger than its children.
// When the parent is less than both children then bubbleDown is complete.
func (h *MinHeap[T]) bubbleDown() {
	parentIdx := 0
	leftIdx, rightIdx := 2*parentIdx+1, 2*parentIdx+2

	for leftIdx < h.Len() {
		parent, leftChild := h.elements[parentIdx], h.elements[leftIdx]

		var rightChild T
		if rightIdx < h.Len() {
			rightChild = h.elements[rightIdx]
		}

		var swapIdx int
		if rightIdx >= h.Len() { // rightChild is nil
			swapIdx = leftIdx
			if h.cmp(parent, leftChild) {
				return
			}
		} else if h.cmp(rightChild, parent) || h.cmp(leftChild, parent) {
			swapIdx = leftIdx
			if h.cmp(rightChild, leftChild) {
				swapIdx = rightIdx
			}
		} else {
			return
		}

		h.swap(parentIdx, swapIdx)
		parentIdx = swapIdx
		leftIdx, rightIdx = 2*parentIdx+1, 2*parentIdx+2
	}
}
