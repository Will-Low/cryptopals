package utils

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHexToBytes(t *testing.T) {
	originalInput := "48656c6c6f"
	bytes, err := HexToBytes("48656c6c6f")
	if err != nil {
		t.Fatal(err)
	}
	convertedBack := hex.EncodeToString(bytes)
	assert.Equal(t, originalInput, convertedBack)
}

func TestBytesToB64(t *testing.T) {
	input := "Hello"
	bytes := []byte(input)
	output := BytesToB64(bytes)
	expected := "SGVsbG8="
	assert.Equal(t, output, expected)
}
