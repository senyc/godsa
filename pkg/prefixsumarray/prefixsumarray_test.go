package prefixsumarray_test

import (
	"github.com/senyc/godsa/pkg/prefixsumarray"
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	test := []struct {
		input          []int
		expectedOutput []int
	}{
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 3, 6, 10, 15}},
	}

	for _, i := range test {
		if output := prefixsumarray.Generate(i.input); !reflect.DeepEqual(output, i.expectedOutput) {
			t.Error(output, "does not equal expected output: ", i.expectedOutput)
		}
	}
}
