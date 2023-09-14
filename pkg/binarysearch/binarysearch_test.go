package binarysearch_test

import (
	"github.com/senyc/godsa/pkg/binarysearch"
	"testing"
)

func TestInt(t *testing.T) {
	test := []struct {
		name  string
		input struct {
			slice  []int
			target int
		}
		output int
	}{
		{"value exists in slice", struct {
			slice  []int
			target int
		}{[]int{1, 2, 3, 4}, 3}, 2},
		{"value does not exist in slice", struct {
			slice  []int
			target int
		}{[]int{1, 2, 3, 4}, 7}, -1},
	}

	for _, v := range test {
		if binarysearch.BinarySearch(v.input.slice, v.input.target) != v.output {
			t.Error("failed:", v.name, "input:", v.input, "does not equal", v.output)
		}
	}
}
