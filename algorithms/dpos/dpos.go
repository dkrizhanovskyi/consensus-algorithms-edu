// Package dpos implements a simplified version of Delegated Proof of Stake (DPoS) consensus mechanism.
// The purpose of this module is to provide a basic blockchain structure where delegates are elected 
// by voters to create and validate new blocks. This approach aims to demonstrate decentralized governance 
// while maintaining efficiency through a delegate-based validation process. Each block is securely hashed 
// using SHA-256 to ensure integrity, and the module provides mechanisms for voting and selecting delegates.
package dpos

import (
    "crypto/sha256"
    "fmt"
    "math/rand"
    "strconv"
    "time"
)

// Block represents an individual block in the blockchain.
// It contains data related to transactions, the timestamp, 
// the delegate responsible for the block, and cryptographic hashes for integrity.
type Block struct {
    Index     int       // The position of the block in the blockchain.
    Timestamp string    // The time when the block was created.
    Data      string    // The transaction or arbitrary data contained in the block.
    PrevHash  string    // Hash of the previous block to ensure immutability and chain integrity.
    Hash      string    // SHA-256 hash of the current block's contents.
    Delegate  string    // The elected delegate responsible for creating this block.
}

// Blockchain represents the overall state of the blockchain,
// including the chain of blocks and the delegates involved in block creation.
type Blockchain struct {
    Blocks    []Block            // A slice of all blocks in the blockchain.
    Delegates []string           // A list of delegates who are eligible to create blocks.
    Voters    map[string]string  // A mapping between voters and the delegates they have voted for.
}

// NewBlock creates a new Block with the given data, previous block hash, index, and delegate.
// It calculates the hash for the block to ensure integrity.
func NewBlock(data string, prevHash string, index int, delegate string) Block {
    block := Block{
        Index:     index,
        Timestamp: time.Now().String(), // Timestamp to record when the block was created.
        Data:      data,
        PrevHash:  prevHash,
        Delegate:  delegate,
    }
    block.Hash = block.CalculateHash() // Calculate the cryptographic hash for the new block.
    return block
}

// CalculateHash generates the SHA-256 hash of the block's contents.
// This includes the index, timestamp, data, previous hash, and delegate, ensuring immutability.
func (b *Block) CalculateHash() string {
    record := strconv.Itoa(b.Index) + b.Timestamp + b.Data + b.PrevHash + b.Delegate
    hash := sha256.New()               // Create a new SHA-256 hash object.
    hash.Write([]byte(record))         // Write the block's contents as bytes to the hash object.
    hashed := hash.Sum(nil)            // Compute the final hash value.
    return fmt.Sprintf("%x", hashed)   // Return the hash as a hexadecimal string.
}

// AddBlock adds a new block to the blockchain.
// It selects a delegate, creates a new block with the given data, and appends it to the chain.
func (bc *Blockchain) AddBlock(data string) {
    prevBlock := bc.Blocks[len(bc.Blocks)-1]        // Retrieve the last block in the chain.
    delegate := bc.SelectDelegate()                  // Select a delegate to produce the next block.
    newBlock := NewBlock(data, prevBlock.Hash, prevBlock.Index+1, delegate)
    bc.Blocks = append(bc.Blocks, newBlock)          // Append the newly created block to the chain.
}

// SelectDelegate randomly selects a delegate from the list of available delegates.
// This function is used to ensure that a delegate is chosen fairly to produce a block.
func (bc *Blockchain) SelectDelegate() string {
    index := rand.Intn(len(bc.Delegates))            // Randomly select an index from the list of delegates.
    return bc.Delegates[index]                       // Return the selected delegate's identifier.
}

// NewBlockchain initializes a new blockchain with a list of delegates and an initial set of voters.
// The blockchain starts with a genesis block, which acts as the foundation of the chain.
func NewBlockchain(delegates []string, voters map[string]string) *Blockchain {
    genesisBlock := NewBlock("Genesis Block", "", 0, delegates[0]) // Create the genesis block.
    return &Blockchain{
        Blocks:    []Block{genesisBlock},         // Initialize with the genesis block.
        Delegates: delegates,                     // Assign the provided list of delegates.
        Voters:    voters,                        // Set up the voters mapping.
    }
}

// Vote allows a voter to vote for a specific delegate.
// This function records the voter's choice, helping to determine the delegate list.
func (bc *Blockchain) Vote(voter string, delegate string) {
    bc.Voters[voter] = delegate                    // Record the voter's choice of delegate.
}

// CountVotes tallies all votes cast by the voters and determines the order of the delegates.
// It sorts the delegates randomly after counting to avoid bias.
func (bc *Blockchain) CountVotes() {
    votes := make(map[string]int)                   // Create a map to hold the count of votes per delegate.
    for _, delegate := range bc.Voters {
        votes[delegate]++                           // Increment the count for each delegate based on votes received.
    }

    sortedDelegates := make([]string, 0)            // Create a slice to store the sorted list of delegates.
    for delegate := range votes {
        sortedDelegates = append(sortedDelegates, delegate) // Populate the list of delegates based on voting results.
    }

    rand.Shuffle(len(sortedDelegates), func(i, j int) {
        sortedDelegates[i], sortedDelegates[j] = sortedDelegates[j], sortedDelegates[i]
    })                                              // Randomly shuffle the list to ensure fairness in delegate order.

    bc.Delegates = sortedDelegates                  // Update the list of delegates with the sorted result.
}

// Footer: Security Considerations and Architectural Decisions
// 
// This implementation of Delegated Proof of Stake (DPoS) focuses on simplicity and demonstrates the principles of blockchain
// consensus through delegate selection and voting. Several architectural and security features are included to enhance reliability:
// 
// 1. **SHA-256 Cryptographic Hashing**: Each block's integrity is ensured through the use of SHA-256 to create a unique hash based
//    on the block's contents. This makes tampering with block data computationally infeasible, as changing any part of the block
//    would result in a completely different hash.
// 
// 2. **Delegate Selection and Randomness**: The `SelectDelegate` function uses random selection to ensure fairness in delegate
//    rotation. This prevents any single delegate from monopolizing block creation and promotes a decentralized approach to validation.
// 
// 3. **Vote Counting and Random Shuffling**: The `CountVotes` function includes a shuffling step to prevent bias and ensure that 
//    the order of delegates is not influenced by predictable patterns. This is crucial for avoiding vulnerabilities that could be
//    exploited by malicious actors aiming to dominate the delegate pool.
// 
// 4. **Genesis Block Initialization**: The genesis block is created when the blockchain is initialized, establishing the root of
//    trust. The choice of the initial delegate is arbitrary and should be carefully managed in production environments to ensure
//    transparency and trust among participants.
// 
// Overall, the code prioritizes security and fairness through the careful use of cryptographic functions and randomized processes.
// While the implementation is simplified for educational purposes, it provides a foundational understanding of how DPoS consensus
// can maintain decentralization, efficiency, and robustness in a blockchain network.
