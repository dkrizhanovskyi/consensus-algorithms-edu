// Package pow implements a basic version of the Proof of Work (PoW) consensus algorithm.
// Proof of Work is used to secure the blockchain by requiring participants to solve a computationally intensive puzzle 
// before adding new blocks. The concept of mining, difficulty, and nonce adjustment forms the core of PoW, which was
// popularized by Bitcoin. This package demonstrates the mining process and how blocks are linked together to maintain
// a tamper-proof blockchain.
package pow

import (
    "crypto/sha256"
    "fmt"
    "strconv"
    "time"
)

// Block represents an individual block in the blockchain.
// It contains crucial information like index, timestamp, data, cryptographic hashes, and a nonce value used for mining.
type Block struct {
    Index     int    // Position of the block in the blockchain.
    Timestamp string // The time when the block was created.
    Data      string // The transaction or arbitrary data contained within the block.
    PrevHash  string // The hash of the previous block to maintain immutability and chain linkage.
    Hash      string // SHA-256 hash of the current block's contents.
    Nonce     int    // Nonce is the number that miners adjust to find a valid hash under the set difficulty.
}

// Blockchain represents the distributed ledger that consists of a chain of blocks.
// Blocks are mined and added to this chain, ensuring that every block is valid and consistent with previous ones.
type Blockchain struct {
    Blocks []Block // A slice containing all blocks in the blockchain.
}

// NewBlock creates a new block, initializes it with given data, and mines it to ensure it meets the difficulty criteria.
// Mining involves adjusting the nonce until a hash with the correct number of leading zeros is found.
func NewBlock(data string, prevHash string, index int) Block {
    block := Block{
        Index:     index,
        Timestamp: time.Now().String(), // Record the time when the block is created.
        Data:      data,
        PrevHash:  prevHash,
        Nonce:     0, // Initialize nonce to zero, which will be incremented during mining.
    }
    block.MineBlock() // Mine the block to find a valid hash that meets the difficulty requirement.
    return block
}

// CalculateHash generates a SHA-256 hash of the block's contents.
// The hash includes the block's index, timestamp, data, previous hash, and nonce.
func (b *Block) CalculateHash() string {
    record := strconv.Itoa(b.Index) + b.Timestamp + b.Data + b.PrevHash + strconv.Itoa(b.Nonce)
    hash := sha256.New()                // Create a new SHA-256 hash object.
    hash.Write([]byte(record))          // Write the concatenated block data to the hash.
    hashed := hash.Sum(nil)             // Compute the hash.
    return fmt.Sprintf("%x", hashed)    // Return the hash as a hexadecimal string.
}

// MineBlock performs the Proof of Work mining process to find a valid hash for the block.
// The mining difficulty is represented by the number of leading zeros in the hash.
func (b *Block) MineBlock() {
    difficulty := 4                     // Set the mining difficulty; 4 leading zeros are required.
    target := "0000"                    // Target pattern that the hash must match (difficulty level of 4).
    
    // Increment the nonce and recalculate the hash until the hash has the required number of leading zeros.
    for b.Hash[:difficulty] != target {
        b.Nonce++                       // Increment nonce to generate a new hash.
        b.Hash = b.CalculateHash()      // Calculate the new hash with the updated nonce.
    }
    // Once the valid hash is found, the block is ready to be added to the blockchain.
}

// AddBlock creates a new block with the given data, mines it, and appends it to the blockchain.
func (bc *Blockchain) AddBlock(data string) {
    prevBlock := bc.Blocks[len(bc.Blocks)-1]         // Retrieve the last block in the chain.
    newBlock := NewBlock(data, prevBlock.Hash, prevBlock.Index+1) // Create a new block based on the previous block.
    bc.Blocks = append(bc.Blocks, newBlock)          // Append the newly mined block to the blockchain.
}

// NewBlockchain initializes a new blockchain with a genesis block.
// The genesis block serves as the first block in the blockchain, establishing the foundation of the chain.
func NewBlockchain() *Blockchain {
    genesisBlock := NewBlock("Genesis Block", "", 0) // Create the genesis block (index 0).
    return &Blockchain{[]Block{genesisBlock}}        // Initialize blockchain with the genesis block.
}

// Footer: Security Considerations and Architectural Decisions
//
// This implementation of Proof of Work (PoW) consensus demonstrates the essential principles of mining and achieving consensus
// in a distributed blockchain environment. Below are the key architectural decisions and security considerations:
//
// 1. **Cryptographic Hashing**: SHA-256 is used to hash the block contents, including the nonce, to produce a unique identifier for each block.
//    This ensures data integrity—if any part of the block is altered, the resulting hash will be completely different,
//    making tampering immediately evident.
//
// 2. **Proof of Work Mining**: The mining process involves incrementing the nonce until a hash is found that matches the required
//    difficulty. This computational challenge ensures that adding a new block is resource-intensive, which deters malicious actors
//    from attempting to alter the blockchain, as they would need to re-mine all subsequent blocks.
//
// 3. **Mining Difficulty**: The difficulty level is set to 4 leading zeros for simplicity. In real-world systems, this value is adjusted dynamically
//    to control the rate of block creation. A higher difficulty makes mining more challenging, thus increasing the security of the network.
//    In production environments, adjusting difficulty helps maintain a consistent rate of block generation.
//
// 4. **Tamper Resistance**: The hash of each block includes the hash of the previous block, creating a linked chain. This ensures that
//    any change to one block would require re-mining all subsequent blocks, making it computationally infeasible to tamper with
//    the blockchain without controlling the majority of computational resources (as seen in the "51% attack" scenario).
//
// 5. **Genesis Block**: The genesis block is created when the blockchain is initialized and serves as the root of trust for the chain.
//    The integrity of the entire chain depends on the immutability of this initial block.
//
// This basic implementation highlights how PoW contributes to the security and immutability of blockchain networks by leveraging 
// computational difficulty. It also provides a foundational understanding of how blockchain data structures can be secured against
// unauthorized changes and how decentralized consensus is achieved through mining.
