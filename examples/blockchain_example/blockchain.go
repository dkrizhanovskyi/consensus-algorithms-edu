// Package main demonstrates the use of the Proof of Work (PoW) consensus algorithm to manage a simple blockchain.
// The Proof of Work consensus algorithm requires miners to solve computational puzzles in order to add new blocks to the blockchain.
// This simulation is designed to illustrate the basic mechanics of how a blockchain works when using PoW for consensus.
// In this example, blocks are mined sequentially, with each new block being appended to the blockchain.
package main

import (
    "fmt"                           // The fmt package is used for formatted I/O, particularly to print output to the console.
    "consensus-algorithms-edu/algorithms/pow" // Import the Proof of Work implementation from the consensus-algorithms-edu module.
)

func main() {
    // Initialize a new blockchain using the Proof of Work algorithm.
    blockchain := pow.NewBlockchain()

    // Add new blocks with specific data to the blockchain.
    blockchain.AddBlock("First block data")
    blockchain.AddBlock("Second block data")
    blockchain.AddBlock("Third block data")

    // Iterate over each block in the blockchain and print the block's details.
    for _, block := range blockchain.Blocks {
        fmt.Printf("Index: %d\nTimestamp: %s\nData: %s\nPrevious Hash: %s\nHash: %s\n\n", 
            block.Index, block.Timestamp, block.Data, block.PrevHash, block.Hash)
    }
}

// Footer: Overview and Execution Flow
//
// In this example, a blockchain is initialized with the Genesis block, which is the foundation of the chain.
// Subsequently, three additional blocks are appended to the chain, each containing some arbitrary data ("First block data", "Second block data", etc.).
// The Proof of Work algorithm is used to mine each block, requiring a computationally intensive process to ensure the integrity of the chain.
//
// Key Steps:
// 1. **Blockchain Initialization**: The blockchain is initialized using `pow.NewBlockchain()`, which creates the Genesis block.
// 2. **Block Addition**: New blocks are added using the `AddBlock()` function, which mines each block before adding it to the blockchain.
// 3. **Block Mining**: Each block requires a valid hash to be found through a Proof of Work computation, which ensures the blockchain's immutability.
// 4. **Block Data Display**: After the blockchain is constructed, the details of each block, such as index, timestamp, data, previous hash, and current hash, are printed.
//
// The primary purpose of this example is to demonstrate how the Proof of Work consensus mechanism ensures that each new block
// added to the blockchain is computationally verified, making the blockchain secure and immutable.
