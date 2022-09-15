package kata1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKata1(t *testing.T) {
	t.Run("description", func(t *testing.T) {
		actual := kata1(2, 2)
		expected := []int{2}

		assert.Equal(t, expected, actual, "There should be one item in the slice")

	})
}

