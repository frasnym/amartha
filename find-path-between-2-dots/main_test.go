package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {

	t.Run("path_is_available", func(t *testing.T) {
		var input Input = Input{
			{0, 0, 0, -1, 0},
			{-1, 0, 0, -1, -1},
			{0, 0, 0, -1, 0},
			{-1, 0, 0, 0, 0},
			{0, 0, -1, 0, 0},
		}

		result := isPathAvailable(input)
		assert.True(t, result)
	})

	t.Run("path_is_not_available", func(t *testing.T) {
		var input Input = Input{
			{-1, 0, 0, -1, -1},
			{0, 0, 0, -1, 0},
			{-1, 0, -1, 0, 0},
			{0, 0, -1, 0, 0},
		}

		result := isPathAvailable(input)
		assert.False(t, result)
	})
}
