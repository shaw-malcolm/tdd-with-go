package template

import "testing"

func TestTemplateFunc(t *testing.T) {
	t.Run("description", func(t *testing.T) {
		got := templateFunc()
		want := "Hello, World"

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	t.Run("description", func(t *testing.T) {
		got := len(templateFunc())
		want := 12

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
