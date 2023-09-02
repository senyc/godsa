package minheap

type MinHeap []int

func (h MinHeap) leftChild(i int) int {
	return 2*i + 1
}

func (h MinHeap) rightChild(i int) int {
	return 2*i + 2
}

func (h *MinHeap) heapify(i, heapSize int) {
	hv := *h

	l := h.leftChild(i)
	r := h.rightChild(i)
	min := i
	if l < heapSize && hv[l] < hv[i] {
		min = l
	}

	if r < heapSize && hv[r] < hv[min] {
		min = r
	}

	if min != i {
		hv[i], hv[min] = hv[min], hv[i]
		h.heapify(min, heapSize)
	}
}

func BuildHeap(a []int) *MinHeap {
	h := MinHeap(a)
	for i := int(len(a)/2) - 1; i >= 0; i-- {
		h.heapify(i, len(a))
	}
	return &h
}
