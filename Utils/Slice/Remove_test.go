package Slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemove(t *testing.T) {
	t.Run("can remove first element", func(t *testing.T) {
		slice := []int{1, 2}

		assert.Equal(t, []int{2}, RemoveFromSlice(1, slice))
	})

	t.Run("can remove second element", func(t *testing.T) {
		slice := []string{"foo", "bar"}

		assert.Equal(t, []string{"foo"}, RemoveFromSlice("bar", slice))
	})

	t.Run("can remove middle element", func(t *testing.T) {
		slice := []string{"a", "b", "c"}

		assert.Equal(t, []string{"a", "c"}, RemoveFromSlice("b", slice))
	})

	t.Run("can remove multiple elements", func(t *testing.T) {
		slice := []int{1, 1, 2, 1, 3, 1}

		assert.Equal(t, []int{2, 3}, RemoveFromSlice(1, slice))
	})

	t.Run("creates independent slice", func(t *testing.T) {
		slice := []string{"foo", "bar"}

		new := RemoveFromSlice("foo", slice)
		new[0] = "abc"

		assert.Equal(t, []string{"foo", "bar"}, slice)
	})

	t.Run("is empty after Removing last element", func(t *testing.T) {
		slice := []string{"foo"}

		assert.Equal(t, []string{}, RemoveFromSlice("foo", slice))
	})

	t.Run("Empty slice stays empty", func(t *testing.T) {
		slice := []int{}

		assert.Equal(t, slice, RemoveFromSlice(0, slice))
	})
}
