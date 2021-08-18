package main

import (
	"fmt"
	"log"

	"github.com/Will-Low/cryptopals/utils"
)

func main() {
	input1 := "1c0111001f010100061a024b53535009181c"
	input2 := "686974207468652062756c6c277320657965"
	input1Bytes, err := utils.HexToBytes(input1)
	if err != nil {
		log.Fatal(err)
	}
	input2Bytes, err := utils.HexToBytes(input2)
	if err != nil {
		log.Fatal(err)
	}
	output, err := utils.XorBytes(input1Bytes, input2Bytes)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%x", output)
}
