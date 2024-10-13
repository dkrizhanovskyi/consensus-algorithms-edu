// Package pos implements a simplified version of the Proof of Stake (PoS) consensus algorithm.
// Proof of Stake is a consensus mechanism that selects validators to propose and validate blocks 
// based on the amount of cryptocurrency they "stake" in the network. This approach is designed to 
// be more energy-efficient compared to Proof of Work, as it does not require extensive computational
// effort. In this implementation, validators are chosen probabilistically, weighted by their stake, 
// to append new blocks to the blockchain.
package pos

import (
    "crypto/sha256"
    "fmt"
    "math/rand"
    "strconv"
    "time"
)

// Block represents an individual block in the blockchain.
// It contains critical information such as the block index, timestamp, data, cryptographic hashes,
// and the validator who proposed the block.
type Block struct {
    Index     int    // The position of the block in the blockchain.
    Timestamp string // The time when the block was created.
    Data      string // The transaction or arbitrary data contained in the block.
    PrevHash  string // The hash of the previous block to ensure immutability.
    Hash      string // SHA-256 hash of the current block's contents.
    Validator string // The validator responsible for validating and adding this block.
}

// Blockchain represents the state of the distributed ledger.
// It contains the chain of blocks, a list of validators, and a map of stakes held by validators.
type Blockchain struct {
    Blocks     []Block          // A slice of all blocks in the blockchain.
    Validators []string         // A list of validator nodes eligible to propose blocks.
    Stakes     map[string]int   // A map of validators to their respective stake values.
}

// NewBlock creates a new Block given data, the previous block's hash, the index, and the validator's ID.
// It calculates the cryptographic hash of the block to ensure its integrity.
func NewBlock(data string, prevHash string, index int, validator string) Block {
    block := Block{
        Index:     index,
        Timestamp: time.Now().String(), // Set the block's timestamp to the current time.
        Data:      data,
        PrevHash:  prevHash,
        Validator: validator,
    }
    block.Hash = block.CalculateHash() // Calculate the block's hash for integrity and immutability.
    return block
}

// CalculateHash generates the SHA-256 hash of the block's contents.
// This ensures immutability; any change to the block's contents results in a different hash.
func (b *Block) CalculateHash() string {
    record := strconv.Itoa(b.Index) + b.Timestamp + b.Data + b.PrevHash + b.Validator
    hash := sha256.New()                // Create a new SHA-256 hash object.
    hash.Write([]byte(record))          // Write the concatenated block data to the hash object.
    hashed := hash.Sum(nil)             // Compute the final hash value.
    return fmt.Sprintf("%x", hashed)    // Return the hash as a hexadecimal string.
}

// AddBlock adds a new block to the blockchain.
// It selects a validator based on their stake, creates a new block, and appends it to the blockchain.
func (bc *Blockchain) AddBlock(data string) {
    prevBlock := bc.Blocks[len(bc.Blocks)-1]          // Retrieve the latest block in the blockchain.
    validator := bc.SelectValidator()                 // Select a validator based on their stake.
    newBlock := NewBlock(data, prevBlock.Hash, prevBlock.Index+1, validator) // Create the new block.
    bc.Blocks = append(bc.Blocks, newBlock)           // Append the newly created block to the blockchain.
}

// SelectValidator selects a validator to propose the next block based on the stakes of each validator.
// The probability of selection is directly proportional to the stake value.
func (bc *Blockchain) SelectValidator() string {
    totalStake := 0
    // Calculate the total stake of all validators.
    for _, stake := range bc.Stakes {
        totalStake += stake
    }

    // Pick a random number in the range of [0, totalStake).
    pick := rand.Intn(totalStake)
    runningTotal := 0

    // Iterate through the validators and accumulate their stakes until the random number is within a range.
    for validator, stake := range bc.Stakes {
        runningTotal += stake
        if runningTotal > pick {
            return validator // The validator whose range contains 'pick' is selected.
        }
    }

    return "" // This should never be reached if the logic above is correct.
}

// NewBlockchain initializes a new blockchain with a list of validators and their respective stakes.
// The blockchain starts with a genesis block, which is always the first block in the chain.
func NewBlockchain(validators []string, stakes map[string]int) *Blockchain {
    genesisBlock := NewBlock("Genesis Block", "", 0, validators[0]) // Create the genesis block.
    return &Blockchain{
        Blocks:     []Block{genesisBlock},  // Initialize with the genesis block.
        Validators: validators,             // Assign the provided list of validators.
        Stakes:     stakes,                 // Set up the validators' stakes.
    }
}

// Footer: Security Considerations and Architectural Decisions
//
// This implementation of Proof of Stake (PoS) consensus demonstrates how validators are selected 
// to propose and validate blocks based on the amount of cryptocurrency they have staked in the network.
// PoS aims to make blockchain consensus more energy-efficient compared to Proof of Work (PoW) while
// still providing incentives for honest participation and network security.
//
// 1. **Cryptographic Hashing**: Each block's integrity is ensured using SHA-256 hashing. The hash is calculated
//    from the block's index, timestamp, data, previous hash, and the validator's identity. This ensures that the
//    block is tamper-proof, as changing any component of the block results in an entirely different hash.
//
// 2. **Validator Selection by Stake**: In PoS, validators are selected probabilistically, with higher stakes increasing
//    the likelihood of being chosen. This approach aims to distribute the validation responsibility fairly among
//    participants based on their contribution to the network. The randomness involved ensures that no validator
//    consistently dominates the network, which promotes decentralization.
//
// 3. **Fairness and Decentralization**: To ensure fairness, validators are selected based on their relative stake in the system.
//    This means that the larger a validator's stake, the more likely they are to be selected. However, unlike PoW, where 
//    computational power is the primary determinant, PoS leverages economic incentives to maintain network security.
//
// 4. **Simplified Model**: This code provides a basic version of PoS for educational purposes. In a real-world scenario,
//    additional measures such as slashing (penalizing dishonest behavior), delegation, and complex staking reward mechanisms
//    would be implemented to further enhance security, prevent abuse, and maintain the integrity of the network.
//
// The Proof of Stake model implemented here is intended to illustrate the fundamental concepts of stake-based consensus.
// It captures the basic mechanism of validator selection and block creation while highlighting the efficiency and reduced 
// resource usage that PoS offers compared to computationally intensive PoW approaches.
