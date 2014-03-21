package trie

import "testing"

func TestCheck(t *testing.T) {
	var trie = NewTrie()
	trie.Insert("abc")
	trie.Insert("abc")
	trie.Insert("abc")

	if trie.Check("ef") != -1 {
		t.Errorf("no matched test failed")
	}
	if trie.Check("ab") != 0 {
		t.Error("substring matched failed")
	}
	if trie.Check("abc") != 3 {
		t.Errorf("matched failed")
	}
}

func TestSuffix(t *testing.T) {
	trie := NewTrie()
	trie.Insert("abcde")
	trie.Insert("abcef")
	trie.Insert("abcfgh")
	trie.Insert("abcfgh")

	m := trie.Suffix("abc")

	if f, ok := m["de"]; ok {
		if f != 1 {
			t.Error("de matched frequency is wrong")
		}
	} else {
		t.Error("de not matched")
	}

	if f, ok := m["ef"]; ok {
		if f != 1 {
			t.Error("cef matched frequency is wrong")
		}
	} else {
		t.Error("ef not matched")
	}

	if f, ok := m["fgh"]; ok {
		if f != 2 {
			t.Error("fgh matched frequency is wrong")
		}
	} else {
		t.Error("abcfgh not matched")
	}
}
