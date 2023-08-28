package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToJSON(t *testing.T) {

	type Entry struct {
		Key string `json:"key"`
	}

	tests := []struct {
		name     string
		value    interface{}
		expected string
	}{
		{
			name:     "succeed",
			value:    Entry{Key: "value"},
			expected: "{\"key\":\"value\"}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			actual := ToJSON(tt.value)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestTestToJSON_Failure(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not  panic")
		}
	}()

	value := map[string]interface{}{
		"foo": make(chan int),
	}

	expected := "{\n    \"key\": \"value\"\n}"

	actual := ToJSON(value)
	assert.Equal(t, expected, actual)
}
