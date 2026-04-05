package transform

import (
	"testing"

	"github.com/v1ejo/scramvid/internal/assert"
)

func TestGenerateSeed(t *testing.T) {
	tests := []struct {
		name        string
		key1        string
		key2        string
		expectEqual bool
	}{
		{
			name:        "same key gives same seed",
			key1:        "a",
			key2:        "a",
			expectEqual: true,
		},
		{
			name:        "different keys give different seeds",
			key1:        "a",
			key2:        "b",
			expectEqual: false,
		},
		{
			name:        "empty key",
			key1:        "",
			key2:        "",
			expectEqual: true,
		},
		{
			name:        "similar keys different result",
			key1:        "abc",
			key2:        "abd",
			expectEqual: false,
		},
		{
			name:        "long key stability",
			key1:        "this-is-a-very-long-key-for-testing",
			key2:        "this-is-a-very-long-key-for-testing",
			expectEqual: true,
		},
		{
			name:        "order matters",
			key1:        "ab",
			key2:        "ba",
			expectEqual: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seed1 := generateSeed(tt.key1)
			seed2 := generateSeed(tt.key2)

			if tt.expectEqual {
				assert.Equal(t, seed1, seed2)
			} else {
				assert.NotEqual(t, seed1, seed2)
			}
		})
	}
}
