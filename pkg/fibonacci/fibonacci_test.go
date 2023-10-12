package fibonacci_test

import (
	"reflect"
	"testing"

	"github.com/senyc/godsa/pkg/fibonacci"
)

func TestZeroInput(t *testing.T) {
	_, err := fibonacci.Generate(0)

	if err == nil {
		t.Error("0 input did not raise error")
	}
}

func TestBaseInput(t *testing.T) {
	arr, err := fibonacci.Generate(1)

	if err != nil {
		t.Error("1 input did raise error")
	}

	if len(arr) != 1 || arr[0] != 1 {
		t.Error("1 input did return correct array")
	}

	_, err = fibonacci.Generate(1)

	if err != nil {
		t.Error("1 input did not raise error")
	}

	newArr, err := fibonacci.Generate(2)

	if len(newArr) != 2 || (newArr[0] != 1 && newArr[1] != 1) {
		t.Error("2 input did not return correct array")
	}
}

func TestSmallInput(t *testing.T) {
	var arr []int
	var err error
	test := []struct {
		num      int
		expected []int
	}{
		{5, []int{1, 1, 2, 3, 5}},
		{4, []int{1, 1, 2, 3}},
		{10, []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}},
	}
	for _, i := range test {
		arr, err = fibonacci.Generate(i.num)
		if err != nil {
			t.Error("Input incorrectly led to error")
		}
		if !reflect.DeepEqual(arr, i.expected) {
			t.Error(arr, "not equal to ", i.expected)
		}
	}
}
