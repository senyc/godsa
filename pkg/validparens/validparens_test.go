package validparens_test

import (
	"github.com/senyc/godsa/pkg/validparens"
	"testing"
)

func TestBaseCase(t *testing.T) {
	m := map[string]bool{
		"()()()":   true,
		")()()":    false,
		"{[]()()}": true,
		"({[]()()})": true,
		"({[]()()}]": false,
	}

	for k, v := range m {
		if a := validparens.ValidParens(k); a != v {
			t.Error(k, "is not", a)
		}
	}
}
