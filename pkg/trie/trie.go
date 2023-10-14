package trie

type TrieNode struct {
	char     rune
	children map[rune]*TrieNode
	isWord   bool
}

type Trie struct {
	root *TrieNode
}

func newNode() *TrieNode {
	tn := new(TrieNode)
	tn.children = make(map[rune]*TrieNode)
	return tn
}

func New() *Trie {
	tr := new(Trie)
	tr.root = newNode()
	return tr
}

func (t *Trie) Add(word string) {
	currNode := t.root
	for _, i := range word {
		if val, ok := currNode.children[i]; ok {
			currNode = val
		} else {
			nt := newNode()
			nt.char = i
			currNode.children[i] = nt
			currNode = nt
		}
	}
	currNode.isWord = true
}

func (t *Trie) WordExist(word string) bool {
	currNode := t.root
	if t.root == nil {
		return false
	}
	for _, i := range word {
		if val, ok := currNode.children[i]; ok {
			currNode = val
		} else {
			return false
		}
	}
	return currNode.isWord
}

func countNodes(t *TrieNode) int {
	if len(t.children) == 0 {
		return 0
	}
	count := 1
	for i := range t.children {
		count += countNodes(t.children[i])
	}
	return count
}

func CountSubstrings(s string) int {
	tr := New()
	for i := 0; i < len(s); i++ {
		tr.Add(s[i:])
	}
	return countNodes(tr.root)
}
