package datatypes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	// Test NewSet
	values := []string{"a", "b", "c"}
	set := NewSet(values, func(s string) string { return s })
	assert.Len(t, set, 3)

	// Test Add
	set.Add("d")
	assert.Contains(t, set, "d")

	// Test Contains
	assert.True(t, set.Contains("a"))
	assert.False(t, set.Contains("x"))

	// Test Delete
	set.Delete("b")
	assert.NotContains(t, set, "b")
}
