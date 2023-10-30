package search

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

}

func Test_Keys(t *testing.T) {

}

func Test_KeysWithPrefix(t *testing.T) {

}

func Test_KeysThatMatch(t *testing.T) {

}

func Test_LongestPrefixOf(t *testing.T) {

}
