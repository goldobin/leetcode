package trie

import (
	"github.com/stretchr/testify/assert"
	"math/rand/v2"
	"testing"
)

func TestTrie(t *testing.T) {
	type step func(int, *testing.T, *Trie) bool
	insert := func(s string) step {
		return func(n int, t *testing.T, tr *Trie) bool {
			tr.Insert(s)
			return true
		}
	}

	search := func(s string, want bool) step {
		return func(n int, t *testing.T, tr *Trie) bool {
			return assert.Equalf(t, want, tr.Search(s), "%03d: Search(%s)", n, s)
		}
	}

	startsWith := func(s string, want bool) step {
		return func(n int, t *testing.T, tr *Trie) bool {
			return assert.Equalf(t, want, tr.StartsWith(s), "%03d: StartsWith(%s)", n, s)
		}
	}

	tests := []struct {
		name  string
		steps []step
	}{
		{
			name: "case 01 leetcode",
			steps: []step{
				insert("apple"),
				search("apple", true),
				search("app", false),
				startsWith("app", true),
				insert("app"),
				search("app", true),
			},
		},
		{
			name: "case 02",
			steps: []step{
				search("", false),
				search("app", false),
				startsWith("app", false),
			},
		},
		{
			name: "case 03.1 insert a",
			steps: []step{
				search("a", false),
				insert("a"),
				search("a", true),
				search("apple", false),
				startsWith("app", false),
			},
		},
		{
			name: "case 03.1 insert z",
			steps: []step{
				search("z", false),
				insert("z"),
				search("z", true),
				startsWith("z", true),
				search("zap", false),
				startsWith("zap", false),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Constructor()
			for i, step := range tt.steps {
				n := i + 1
				if !step(n, t, &tr) {
					t.Errorf("step %d failed", n)
					break
				}
			}
		})
	}
}

func TestTrie_LargeInserts(t *testing.T) {

	alfabet := "abcdefghijklmnopqrstuvwxyz"
	tr := Constructor()

	for iteration := 0; iteration < 10000; iteration++ {
		bs := make([]byte, 2000)

		for i := 0; i < len(bs); i++ {
			r := rand.IntN(len(alfabet))
			bs[i] = alfabet[r]
		}

		s := string(bs)

		tr.Insert(s)
		gotSearch := tr.Search(s)
		gotStartsWith := tr.StartsWith(s[:len(s)-1])

		assert.True(t, gotSearch)
		assert.True(t, gotStartsWith)
	}
}

func BenchmarkTrie_Insert(b *testing.B) {
	alfabet := "abcdefghijklmnopqrstuvwxyz"
	tr := Constructor()

	for i := 0; i < b.N; i++ {
		j := rand.IntN(len(alfabet))

		tr.Insert(alfabet[:j])
	}
}

func BenchmarkTrie_Search(b *testing.B) {
	alfabet := "abcdefghijklmnopqrstuvwxyz"
	tr := Constructor()

	for i := 0; i < b.N; i++ {
		j := rand.IntN(len(alfabet))

		tr.Search(alfabet[:j])
	}
}

func BenchmarkTrie_StartsWith(b *testing.B) {
	tr := Constructor()
	s := "abcdefghijklmnopqrstuvwxyz"

	for i := 0; i < b.N; i++ {
		j := rand.IntN(len(s))

		tr.StartsWith(s[:j])
	}
}
