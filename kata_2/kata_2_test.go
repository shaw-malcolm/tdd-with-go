package kata2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKata2(t *testing.T) {
	// test with standard Go syntax
	t.Run("description", func(t *testing.T) {
		got := kata2()
		want := ""

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	// test with Testify
	t.Run("description", func(t *testing.T) {
		actual := len(kata2())
		expected := 12
		assert.Equal(t, expected, actual, "Length of string should be one less than a baker's dozen")
	})
}
