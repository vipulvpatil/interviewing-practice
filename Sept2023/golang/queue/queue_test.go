package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_String(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		s := Queue[string]{}
		s.Enqueue("one")
		s.Enqueue("two")
		s.Enqueue("three")
		popped := s.Dequeue()
		assert.Equal(t, "one", *popped)
		assert.Equal(t, false, s.IsEmpty())
		s.Enqueue("four")
		s.Enqueue("five")
		popped = s.Dequeue()
		assert.Equal(t, "two", *popped)
		popped = s.Dequeue()
		assert.Equal(t, "three", *popped)
		popped = s.Dequeue()
		assert.Equal(t, "four", *popped)
		popped = s.Dequeue()
		assert.Equal(t, "five", *popped)
		assert.Equal(t, true, s.IsEmpty())
		assert.Nil(t, s.Dequeue())
	})
}
