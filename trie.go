package trie

// Trie is trie structure
type Trie struct {
    root *TrieNode
}

// NewTrie create an empty trie
func NewTrie() *Trie {
    trie := &Trie{root: NewTrieNode()}
    return trie
}

// Insert method insert a word into the trie
func (trie *Trie) Insert(word string) {
    trie.root.insert(word)
}

// Check return the word's matched frequency
func (trie *Trie) Check(word string) int {
    node := trie.root.match(word)
    if node == nil {
        return -1
    }
    return node.frequency
}

// Suffix return pattern's suffix string
func (trie *Trie) Suffix(pattern string) map[string]int {
    node := trie.root.match(pattern)
    if node == nil {
        return nil
    }
    return node.Suffix()
}

// Travel return a map contains all saved word and frequency
func (trie *Trie) Travel() map[string]int {
    return trie.root.Suffix()
}
