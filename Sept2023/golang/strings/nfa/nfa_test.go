package nfa

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NFA(t *testing.T) {
	t.Run("test regex is correctly processed using an NFA", func(t *testing.T) {
		nfa := NewNFA("((A*B|AC)D)")
		fmt.Println(nfa.epGraph)
		assert.True(t, nfa.Match("AABD"))
		assert.False(t, nfa.Match("ABDC"))
		assert.True(t, nfa.Match("AAAAAAABD"))
		assert.True(t, nfa.Match("ACD"))
	})
}
