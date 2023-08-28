package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		name     string
		isError  bool
		expected *Configurations
	}{
		{
			name:     "succeed",
			expected: Default,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			actual, err := Load()

			if tt.isError {
				require.Error(t, err)
				require.Nil(t, actual)
			} else {
				require.Nil(t, err)
				require.Equal(t, tt.expected, actual)
			}

		})
	}
}
