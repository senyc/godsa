package gaussianfunction_test

import (
	"github.com/senyc/godsa/pkg/gaussianfunction"
	"testing"
)

func TestSum(t *testing.T) {
	test := []struct {
		name   string
		input  []int
		output int
	}{
		{"Ends at 1", []int{1}, 1},
		{"Ends at 2", []int{1, 2}, 3},
		{"Ends at 5", []int{1, 2, 3, 4, 5}, 15},
	}

	for _, v := range test {
		if result := gaussianfunction.GuassianFunction(v.input); result != v.output {
			t.Error("failed:", v.name, "input:", result, "does not equal", v.output)
		}
	}
}
