package minheap_test

import (
	"fmt"
	"testing"

	"github.com/senyc/godsa/pkg/minheap"
)

func TestBuildHeap(t *testing.T) {
	a := []int{3, 9, 2, 1, 4, 5}
	expectedHeap := minheap.MinHeap{1, 3, 2, 9, 4, 5}

	h := minheap.BuildHeap(a)

	// Check if the resulting heap matches the expected heap
	fmt.Println(*h)
	fmt.Println(expectedHeap)
	for i := range *h {
		if (*h)[i] != expectedHeap[i] {
			t.Errorf("Expected %v at index %d, but got %v", expectedHeap[i], i, (*h)[i])
		}
	}
}
