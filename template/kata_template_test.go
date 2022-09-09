package template

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemplateFunc(t *testing.T) {
	// test with standard Go syntax
	t.Run("description", func(t *testing.T) {
		got := templateFunc()
		want := "Hello, World"

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	// test with Testify
	t.Run("description", func(t *testing.T) {
		actual := len(templateFunc())
		expected := 12
		assert.Equal(t, expected, actual, "Length of string should be one less than a baker's dozen")
	})
}
