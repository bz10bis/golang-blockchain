package main

import (
	"fmt"
	"strconv"

	"github.com/bz10bis/golang-blockchain/blockchain"
)

// MAIN
func main() {
	fmt.Printf("[INIT] Blockchain")
	chain := blockchain.InitBlockChain()
	fmt.Printf("Add blocks")
	chain.AddBlock("First Block")
	chain.AddBlock("Second Block")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := blockchain.NewProof(block)
		fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
