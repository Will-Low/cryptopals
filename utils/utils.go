package utils

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"log"
)

func HexToBytes(hexText string) ([]byte, error) {
	return hex.DecodeString(hexText)
}

func BytesToB64(bytes []byte) string {
	return base64.StdEncoding.EncodeToString(bytes)
}

func XorBytes(input1 []byte, input2 []byte) ([]byte, error) {
	if !byteLengthEqual(input1, input2) {
		return nil, errors.New("Inputs do not have equal length")
	}
	xored := []byte{}
	for i, _ := range input1 {
		xored = append(xored, input1[i]^input2[i])
	}
	return xored, nil
}

func XorHexStrings(input1 string, input2 string) (string, error) {
	if !stringLengthEqual(input1, input2) {
		return "", errors.New("Inputs do not have equal length")
	}

	input1Bytes, err := HexToBytes(input1)
	if err != nil {
		log.Fatal(err)
	}

	input2Bytes, err := HexToBytes(input2)
	if err != nil {
		log.Fatal(err)
	}

	xored := []byte{}
	for i, _ := range input1Bytes {
		xored = append(xored, input1Bytes[i]^input2Bytes[i])
	}
	return hex.EncodeToString(xored), nil
}

func byteLengthEqual(input1 []byte, input2 []byte) bool {
	if len(input1) == len(input2) {
		return true
	} else {
		return false
	}
}

func stringLengthEqual(input1 string, input2 string) bool {
	if len(input1) == len(input2) {
		return true
	} else {
		return false
	}
}
