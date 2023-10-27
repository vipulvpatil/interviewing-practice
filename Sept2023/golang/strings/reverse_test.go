package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Reverse(t *testing.T) {
	t.Run("reverse works for UTF-8 strings", func(t *testing.T) {
		reversed := Reverse("English")
		assert.Equal(t, "hsilgnE", reversed)
		reversed = Reverse("日本語")
		assert.Equal(t, "語本日", reversed)
	})
	t.Run("reverse does not works for non UTF-8 strings", func(t *testing.T) {
		reversed := Reverse("\xbc\x3d\x20")
		assert.Equal(t, "\x20\x3d\xbc", reversed)
	})
}
