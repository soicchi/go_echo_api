package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateExpectedJSON(t *testing.T) {
	expectedString := "Hello World"
	expectedJSON := GenerateExpectedJSON(expectedString)

	assert.Equal(t, "\"Hello World\"\n", expectedJSON)
}
