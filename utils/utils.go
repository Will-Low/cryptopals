package utils

import (
	"encoding/base64"
	"encoding/hex"
)

func HexToBytes(hexText string) ([]byte, error) {
	return hex.DecodeString(hexText)
}

func BytesToB64(bytes []byte) string {
	return base64.StdEncoding.EncodeToString(bytes)
}
