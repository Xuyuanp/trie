package trie

// TrieNode is a node structure of trie
type TrieNode struct {
	frequency int
	children  map[byte]*TrieNode
}

// NewTrieNode return a new empty trie node
func NewTrieNode() *TrieNode {
	node := &TrieNode{frequency: 0, children: make(map[byte]*TrieNode)}
	return node
}

// IsValid return if the node match a word
func (node *TrieNode) IsValid() bool {
	return node != nil && node.frequency > 0
}

// GetFrequency return the patter matched frequency
func (node *TrieNode) GetFrequency() int {
	return node.frequency
}

// Suffix return all suffixes of the current pattern
func (node *TrieNode) Suffix() map[string]int {
	res := make(map[string]int)
	node.suffix("", res)
	return res
}

func (node *TrieNode) suffix(pattern string, m map[string]int) {
	for suf, child := range node.children {
		if child.IsValid() {
			m[pattern+string(suf)] = child.frequency
		}
		child.suffix(pattern+string(suf), m)
	}
}

func (node *TrieNode) insert(pattern string) {
	if len(pattern) == 0 {
		return
	}
	first := pattern[0]
	child, ok := node.children[first]
	if !ok {
		child = NewTrieNode()
		node.children[first] = child
	}
	if len(pattern) == 1 {
		child.frequency++
		return
	}

	child.insert(string(pattern[1:]))
}

func (node *TrieNode) match(pattern string) *TrieNode {
	first := pattern[0]
	child, ok := node.children[first]
	if !ok {
		return nil
	}
	if len(pattern) == 1 {
		return child
	}
	return child.match(string(pattern[1:]))
}
