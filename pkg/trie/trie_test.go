package trie_test

import (
	"github.com/senyc/godsa/pkg/trie"
	"testing"
)

func TestEmptySearch(t *testing.T) {
	tr := trie.New()
	if exists := tr.WordExist("test"); exists {
		t.Error("Search for word in empty trie yielded", exists)
	}
}

func TestFalseSearch(t *testing.T) {
	tr := trie.New()
	tr.Add("test")
	if exists := tr.WordExist("testing"); exists {
		t.Error("Search for nonexistent word in  trie yielded", exists)
	}
	if exists := tr.WordExist("tes"); exists {
		t.Error("Search for nonexistent word in  trie yielded", exists)
	}
	if exists := tr.WordExist("app"); exists {
		t.Error("Search for nonexistent word in  trie yielded", exists)
	}
}

func TestValidSearch(t *testing.T) {
	tr := trie.New()
	tr.Add("test")
	tr.Add("testing")
	if exists := tr.WordExist("test"); ! exists {
		t.Error("Search for word in trie where it exists yielded", exists)
	}
	if exists := tr.WordExist("testing"); ! exists {
		t.Error("Search for word in trie where it exists yielded", exists)
	}

	tr.Add("john")
	if exists := tr.WordExist("john"); ! exists {
		t.Error("Search for word in trie where it exists yielded", exists)
	}
}
