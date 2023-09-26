package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_String(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		s := Stack[string]{}
		s.Push("one")
		s.Push("two")
		s.Push("five")
		popped := s.Pop()
		assert.Equal(t, "five", *popped)
		assert.Equal(t, false, s.IsEmpty())
		s.Push("three")
		s.Push("four")
		popped = s.Pop()
		assert.Equal(t, "four", *popped)
		popped = s.Pop()
		assert.Equal(t, "three", *popped)
		popped = s.Pop()
		assert.Equal(t, "two", *popped)
		popped = s.Pop()
		assert.Equal(t, "one", *popped)
		assert.Equal(t, true, s.IsEmpty())
		assert.Nil(t, s.Pop())
	})
}
