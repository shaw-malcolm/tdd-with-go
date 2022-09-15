package kata2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKata2(t *testing.T) {
	t.Run("description", func(t *testing.T) {
		actual := kata2("a")
		expected := "A"
		assert.Equal(t, expected, actual, "Should return capital A")
	})
}
