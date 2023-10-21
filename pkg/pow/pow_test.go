package pow_test

import (
	"github.com/senyc/godsa/pkg/pow"
	"testing"
)

func TestPow(t *testing.T) {
	test := []struct {
		input struct {
			val int
			pow int
		}
		output int
	}{
		{struct {
			val int
			pow int
		}{2, 10}, 1024},
		{struct {
			val int
			pow int
		}{1, 10}, 1},
		{struct {
			val int
			pow int
		}{5, 2}, 25},
	}

	for _, i := range test {
		if result := pow.Pow(i.input.val, i.input.pow); result != i.output {
			t.Error(result, " Does not equal", i.output)
		}
	}
}
