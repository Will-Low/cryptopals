package main

import (
	"fmt"
	"log"

	"github.com/Will-Low/cryptopals/utils"
)

func main() {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	bytes, err := utils.HexToBytes(input)
	if err != nil {
		log.Fatal(err)
	}
	output := utils.BytesToB64(bytes)
	fmt.Println(output)
}
