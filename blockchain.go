package main

import (
	"fmt"
	"strconv"
	"time"
)

// Block defines the structure of a block in the blockchain
type Block struct {
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
}

// Blockchain represents a series of linked blocks
var Blockchain []Block

// createHash is a simplified function to simulate hashing of block data
func createHash(timestamp, data, prevHash string) string {
	return fmt.Sprintf("%x", timestamp+data+prevHash)
}

// NewBlock creates a new Block and returns it
func NewBlock(data string, prevHash string) Block {
	block := Block{
		Timestamp: time.Now().String(),
		Data:      data,
		PrevHash:  prevHash,
	}
	block.Hash = createHash(block.Timestamp, block.Data, block.PrevHash)
	return block
}

// AddBlock adds a new block to the Blockchain
func AddBlock(data string) {
	var prevHash string
	if len(Blockchain) > 0 {
		prevHash = Blockchain[len(Blockchain)-1].Hash
	}
	Blockchain = append(Blockchain, NewBlock(data, prevHash))
}

// ModifyBlock changes the data of a given block at index
func ModifyBlock(index int, newData string) bool {
	if index < 0 || index >= len(Blockchain) {
		return false // Return false if index is out of bounds
	}
	Blockchain[index].Data = newData
	Blockchain[index].Hash = createHash(Blockchain[index].Timestamp, newData, Blockchain[index].PrevHash)
	return true
}

// DisplayAllBlocks prints the details of all blocks in the blockchain
func DisplayAllBlocks() {
	for i, block := range Blockchain {
		fmt.Printf("Block %d: %+v\n", i, block)
	}
}

func main() {
	// Add some blocks to the blockchain
	AddBlock("First Block")
	AddBlock("Second Block")

	// Display all blocks
	fmt.Println("Initially, the blockchain contains:")
	DisplayAllBlocks()

	// Modify a block
	if ModifyBlock(1, "Modified Second Block") {
		fmt.Println("\nAfter modifying the second block:")
		DisplayAllBlocks()
	} else {
		fmt.Println("Failed to modify block.")
	}
}
