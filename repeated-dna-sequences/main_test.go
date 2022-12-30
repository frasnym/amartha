package dna

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckDNAString(t *testing.T) {
	t.Run("should_return_2", func(t *testing.T) {
		result := checkDNAString(4, "CCATATGTATGGATAT")
		assert.Len(t, result, 2)
	})

	t.Run("should_return_1", func(t *testing.T) {
		result := checkDNAString(6, "AGAGCTCCAGAGCTTG")
		assert.Len(t, result, 1)
	})
}

func TestCheckDNAString2(t *testing.T) {
	t.Run("should_return_2", func(t *testing.T) {
		result := checkDNAString2(4, "CCATATGTATGGATAT")
		assert.Len(t, result, 2)
	})

	t.Run("should_return_1", func(t *testing.T) {
		result := checkDNAString2(6, "AGAGCTCCAGAGCTTG")
		assert.Len(t, result, 1)
	})
}
