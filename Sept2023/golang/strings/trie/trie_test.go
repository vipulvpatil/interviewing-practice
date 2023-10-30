package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PutAndGet(t *testing.T) {
	tr := Trie{}
	tr.Put("key1", "key1")
	tr.Put("key2", "key2")
	tr.Put("random_key", "random_key")
	tr.Put("ke", "ke")
	tr.Put("k", "k")
	assert.Equal(t, "key1", tr.Get("key1"))
	assert.Equal(t, "key2", tr.Get("key2"))
	assert.Equal(t, "random_key", tr.Get("random_key"))
	assert.Equal(t, "ke", tr.Get("ke"))
	assert.Equal(t, "k", tr.Get("k"))
	assert.Equal(t, "", tr.Get("key3"))
}

func Test_Delete(t *testing.T) {
	tr := Trie{}
	tr.Put("key1", "key1")
	tr.Put("key2", "key2")
	tr.Put("random_key", "random_key")
	tr.Put("ke", "ke")
	tr.Put("k", "k")
	tr.Delete("key2")
	assert.Equal(t, "key1", tr.Get("key1"))
	assert.Equal(t, "", tr.Get("key2"))
	assert.Equal(t, "random_key", tr.Get("random_key"))
	assert.Equal(t, "ke", tr.Get("ke"))
	assert.Equal(t, "k", tr.Get("k"))
	assert.Equal(t, "", tr.Get("key3"))

}

func Test_Keys(t *testing.T) {
	tr := Trie{}
	tr.Put("key1", "key1")
	tr.Put("key2", "key2")
	tr.Put("key15", "key15")
	tr.Put("random_key", "random_key")
	tr.Put("ke", "ke")
	tr.Put("k", "k")
	assert.ElementsMatch(t, []string{"random_key", "k", "ke", "key1", "key2", "key15"}, tr.Keys())
}

func Test_KeysWithPrefix(t *testing.T) {
	tr := Trie{}
	tr.Put("key1", "key1")
	tr.Put("key2", "key2")
	tr.Put("key15", "key15")
	tr.Put("random_key", "random_key")
	tr.Put("ke", "ke")
	tr.Put("k", "k")
	assert.ElementsMatch(t, []string{"ke", "key1", "key2", "key15"}, tr.KeysWithPrefix("ke"))
}

func Test_KeysThatMatch(t *testing.T) {
	tr := Trie{}
	tr.Put("key1", "key1")
	tr.Put("key2", "key2")
	tr.Put("key15", "key15")
	tr.Put("random_key", "random_key")
	tr.Put("ke", "ke")
	tr.Put("k", "k")
	tr.Put("bat", "bat")
	tr.Put("cat", "cat")
	tr.Put("hat", "hat")
	tr.Put("that", "that")
	tr.Put("this", "this")
	tr.Put("them", "them")
	tr.Put("thor", "thor")
	assert.ElementsMatch(t, []string{"ke"}, tr.KeysThatMatch("ke"))
	assert.ElementsMatch(t, []string{"key1", "key2"}, tr.KeysThatMatch("k.y."))
	assert.ElementsMatch(t, []string{"key1", "key2"}, tr.KeysThatMatch("key."))
	assert.ElementsMatch(t, []string{"thor", "that", "this", "them"}, tr.KeysThatMatch("th.."))
}

func Test_LongestPrefixOf(t *testing.T) {
	tr := Trie{}
	tr.Put("flower", "flower")
	tr.Put("flowery", "flowery")
	tr.Put("flow", "flow")
	tr.Put("flo", "flo")
	tr.Put("flowing", "flowing")
	tr.Put("glow", "glow")
	assert.Equal(t, "flower", tr.LongestPrefixOf("flower"))
	assert.Equal(t, "flo", tr.LongestPrefixOf("float"))
	assert.Equal(t, "flower", tr.LongestPrefixOf("flowering"))
	assert.Equal(t, "", tr.LongestPrefixOf("gloat"))
}
