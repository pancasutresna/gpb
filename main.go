package main

import (
	"fmt"
	"strconv"
)

func main() {
	bc := NewBlockchain()
	bc.AddBlock("Example : Sending 1 btc to Bob")
	bc.AddBlock("Example : Sending 2 more btc to Bob")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash : %x\n", block.PrevBlockHash)
		fmt.Printf("Data : %s\n", block.Data)
		fmt.Printf("Hash : %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
