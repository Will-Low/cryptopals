package utils

import (
	"encoding/hex"
	"testing"

	"github.com/Will-Low/cryptopals/utils"
	"github.com/stretchr/testify/assert"
)

func TestHexToBytes(t *testing.T) {
	originalInput := "48656c6c6f"
	bytes, err := utils.HexToBytes("48656c6c6f")
	if err != nil {
		t.Fatal(err)
	}
	convertedBack := hex.EncodeToString(bytes)
	assert.Equal(t, originalInput, convertedBack)
}

func TestBytesToB64(t *testing.T) {
	input := "Hello"
	bytes := []byte(input)
	output := utils.BytesToB64(bytes)
	expected := "SGVsbG8="
	assert.Equal(t, output, expected)
}

func TestXorBytes(t *testing.T) {
	input1, err := utils.HexToBytes("1c0111001f010100061a024b53535009181c")
	if err != nil {
		t.Fatal(err)
	}

	input2, err := utils.HexToBytes("686974207468652062756c6c277320657965")
	if err != nil {
		t.Fatal(err)
	}

	expectedOutput, err := utils.HexToBytes("746865206b696420646f6e277420706c6179")
	if err != nil {
		t.Fatal(err)
	}

	xored, err := utils.XorBytes(input1, input2)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, xored, expectedOutput)
}

func TestXorHexStrings(t *testing.T) {
	input1 := "1c0111001f010100061a024b53535009181c"
	input2 := "686974207468652062756c6c277320657965"

	output, err := utils.XorHexStrings(input1, input2)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, output, "746865206b696420646f6e277420706c6179")
}
