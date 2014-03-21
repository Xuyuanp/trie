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

func TestTravel(t *testing.T) {
	trie := NewTrie()
	trie.Insert("abcd")
	trie.Insert("abcd")
	trie.Insert("ab")
	trie.Insert("xyz")

	m := trie.Travel()

	if freq, ok := m["ab"]; ok {
		if freq != 1 {
			t.Error("ab matched frequency is wrong")
		}
	} else {
		t.Error("ab not matched")
	}

	if freq, ok := m["abcd"]; ok {
		if freq != 2 {
			t.Error("abcd matched frequency is wrong")
		}
	} else {
		t.Error("abcd not matched")
	}

	if freq, ok := m["xyz"]; ok {
		if freq != 1 {
			t.Error("xyz matched frequency is wrong")
		}
	} else {
		t.Error("xyz not matched")
	}
}
