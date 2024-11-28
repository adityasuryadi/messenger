package test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoginController(t *testing.T) {
	t.Run("Test Login Success", func(t *testing.T) {
		require.Equal(t, 1, 1)
	})

	t.Run("Error when are fields empty", func(t *testing.T) {
		require.Equal(t, 1, 2)
	})
}
