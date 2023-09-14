package reverse_test

import (
	"github.com/senyc/godsa/pkg/reverse"
	"reflect"
	"testing"
)

func TestReverseInt(t *testing.T) {
	test := []struct {
		name   string
		input  []int
		output []int
	}{
		{"sequential test", []int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{"reverse sequential test", []int{4, 3, 2, 1}, []int{1, 2, 3, 4}},
		{"random order test", []int{11, 24, 3, 41}, []int{41, 3, 24, 11}},
	}

	for _, v := range test {
		if !reflect.DeepEqual(reverse.Reverse(v.input), v.output) {
			t.Error("failed:", v.name, "input:", v.input, "does not equal", v.output)
		}
	}
}

func TestReverseString(t *testing.T) {
	test := []struct {
		name   string
		input  []string
		output []string
	}{
		{"random string test", []string{"hi", "hello", "goodbye", "afternoon"}, []string{"afternoon",  "goodbye",  "hello", "hi"}},
	}

	for _, v := range test {
		if !reflect.DeepEqual(reverse.Reverse(v.input), v.output) {
			t.Error("failed:", v.name, "input:", v.input, "does not equal", v.output)
		}
	}
}

func TestReverseByte(t *testing.T) {
	test := []struct {
		name   string
		input  []byte
		output []byte
	}{
		{"sequential order test", []byte{'1', '2', '3', '4'}, []byte{'4', '3', '2', '1'}},
	}

	for _, v := range test {
		if !reflect.DeepEqual(reverse.Reverse(v.input), v.output) {
			t.Error("failed:", v.name, "input:", v.input, "does not equal", v.output)
		}
	}
}
