package storage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShorten(t *testing.T) {
	t.Run("returns analphanumeric short identifier", func(t *testing.T) {
		type testCase struct {
			id       uint32
			expected string
		}
		testCases := []testCase{
			{
				id:       1003,
				expected: "t4",
			},
			{
				id:       0,
				expected: "",
			},
		}

		for _, tc := range testCases {
			actual := ShortenURL(tc.id)
			assert.Equal(t, tc.expected, actual)
		}
	})
	t.Run("is idempotent", func(t *testing.T) {
		for i := 0; i < 100; i++ {
			assert.Equal(t, "us", ShortenURL(1024))
		}
	})
}
