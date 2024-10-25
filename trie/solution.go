package trie

const alfabetSize = 'z' - 'a' + 1

type (
	Trie struct {
		children [alfabetSize]*Trie
		hasValue bool
	}
)

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(s string) {
	if len(s) == 0 {
		t.hasValue = true
		return
	}
	ch := s[0]
	ix := index(ch)
	v := t.children[ix]
	if v == nil {
		v = &Trie{}
		t.children[ix] = v
	}

	v.Insert(s[1:])
}

func (t *Trie) Search(s string) bool {
	if len(s) == 0 {
		return t.hasValue
	}

	v := t.child(s[0])
	if v == nil {
		return false
	}

	return v.Search(s[1:])
}

func (t *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return true
	}

	v := t.child(prefix[0])
	if v == nil {
		return false
	}

	return v.StartsWith(prefix[1:])
}

func (t *Trie) child(ch byte) *Trie {
	return t.children[index(ch)]
}

func index(ch byte) int {
	return int(ch - 'a')
}
