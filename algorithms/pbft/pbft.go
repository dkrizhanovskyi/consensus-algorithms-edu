// Package pbft implements a basic version of the Practical Byzantine Fault Tolerance (PBFT) consensus algorithm.
// PBFT is designed to allow a group of nodes to reach consensus in a distributed system, even in the presence
// of Byzantine faults—nodes that may behave arbitrarily or maliciously. This module simulates a blockchain 
// network using PBFT to demonstrate how nodes work together to validate and commit new blocks, ensuring consistency 
// and resilience against malicious actors. The algorithm requires agreement from at least 2/3 of the nodes 
// for a block to be committed.
package pbft

import (
    "crypto/sha256"
    "fmt"
    "strconv"
    "time"
)

// Block represents an individual block in the blockchain.
// Each block contains essential information, including metadata, the cryptographic hash, and the data it stores.
type Block struct {
    Index     int    // The position of the block in the blockchain.
    Timestamp string // The time when the block was created.
    Data      string // The data contained in the block (e.g., transactions).
    PrevHash  string // The hash of the previous block, establishing continuity of the chain.
    Hash      string // The cryptographic hash of the current block's contents.
}

// Blockchain represents the distributed ledger, which is maintained by nodes.
// It contains an ordered list of blocks, each of which is linked to its predecessor by cryptographic hash.
type Blockchain struct {
    Blocks []Block // A slice of all blocks in the blockchain.
    Nodes  []Node  // A slice representing all nodes participating in PBFT consensus.
}

// Node represents an individual node participating in the PBFT protocol.
// Each node can propose, verify, and commit blocks, and maintains its own state and reference to the blockchain.
type Node struct {
    ID          int         // A unique identifier for the node.
    IsPrimary   bool        // A flag indicating if this node is the primary node (leader).
    State       string      // The state of the node (optional, for future implementation).
    Blockchain  *Blockchain // Reference to the blockchain managed by the node.
}

// NewBlock creates a new block given the data, index, and previous block hash.
// It calculates the hash for the new block to ensure data integrity.
func NewBlock(data string, prevHash string, index int) Block {
    block := Block{
        Index:     index,
        Timestamp: time.Now().String(), // Set the block's timestamp to the current time.
        Data:      data,
        PrevHash:  prevHash,
    }
    block.Hash = block.CalculateHash() // Calculate the block's hash for integrity.
    return block
}

// CalculateHash generates the SHA-256 hash of the block's contents.
// This ensures that each block is uniquely represented and immutable.
func (b *Block) CalculateHash() string {
    record := strconv.Itoa(b.Index) + b.Timestamp + b.Data + b.PrevHash
    hash := sha256.New()                // Create a new SHA-256 hash object.
    hash.Write([]byte(record))          // Write the concatenated block data to the hash.
    hashed := hash.Sum(nil)             // Compute the hash.
    return fmt.Sprintf("%x", hashed)    // Return the hash as a hexadecimal string.
}

// AddBlock appends a new block to the blockchain.
func (bc *Blockchain) AddBlock(block Block) {
    bc.Blocks = append(bc.Blocks, block) // Append the new block to the blockchain.
}

// NewBlockchain initializes a new blockchain with a genesis block, which serves as the root of the chain.
func NewBlockchain() *Blockchain {
    genesisBlock := NewBlock("Genesis Block", "", 0) // Create the genesis block.
    return &Blockchain{
        Blocks: []Block{genesisBlock}, // Initialize with the genesis block.
        Nodes:  []Node{},              // Initialize an empty list of nodes.
    }
}

// ProposeBlock allows the primary node to create a new block proposal.
// It retrieves the latest block and proposes a new block with the given data.
func (n *Node) ProposeBlock(data string) Block {
    prevBlock := n.Blockchain.Blocks[len(n.Blockchain.Blocks)-1] // Get the last block in the chain.
    newBlock := NewBlock(data, prevBlock.Hash, prevBlock.Index+1) // Create a new block based on the latest block.
    return newBlock
}

// BroadcastBlock broadcasts a proposed block to all nodes in the network for verification.
// A block is considered valid if at least 2/3 of nodes approve it.
func (bc *Blockchain) BroadcastBlock(block Block) bool {
    approvals := 0
    totalNodes := len(bc.Nodes)
    
    for _, node := range bc.Nodes {
        if node.VerifyBlock(block) {
            approvals++ // Increment the count if the node approves the block.
        }
    }
    
    // Return true if 2/3 or more nodes approve the block.
    return approvals >= (2 * totalNodes / 3)
}

// VerifyBlock allows a node to verify the validity of a proposed block.
// The node checks if the block's previous hash matches the last block in the chain and if the block hash is valid.
func (n *Node) VerifyBlock(block Block) bool {
    prevBlock := n.Blockchain.Blocks[len(n.Blockchain.Blocks)-1] // Retrieve the latest block in the chain.
    // Verify if the proposed block's previous hash matches the latest block's hash and if the block hash is valid.
    if block.PrevHash == prevBlock.Hash {
        return block.Hash == block.CalculateHash()
    }
    return false
}

// CommitBlock adds a block to the blockchain, once it has been verified and approved by the network.
func (n *Node) CommitBlock(block Block) {
    n.Blockchain.AddBlock(block) // Append the verified block to the blockchain.
}

// RunPBFT initiates the Practical Byzantine Fault Tolerance consensus process.
// The primary node proposes a new block, and if it receives approval from 2/3 of nodes, all nodes commit the block.
func (bc *Blockchain) RunPBFT(data string) {
    primary := bc.Nodes[0]                   // The first node is treated as the primary node (leader).
    newBlock := primary.ProposeBlock(data)   // Primary node proposes a new block.

    // Broadcast the proposed block for verification, and if approved, commit it across all nodes.
    if bc.BroadcastBlock(newBlock) {
        for _, node := range bc.Nodes {
            node.CommitBlock(newBlock)       // Each node commits the approved block.
        }
    }
}

// NewNode creates a new node with the given ID, assigns it as primary or follower, and links it to the blockchain.
func NewNode(id int, isPrimary bool, blockchain *Blockchain) *Node {
    return &Node{
        ID:         id,
        IsPrimary:  isPrimary,
        Blockchain: blockchain,
    }
}

// NewPBFTNetwork initializes a PBFT network with a specified number of nodes.
// The first node is assigned as the primary node, and all nodes are linked to the blockchain.
func NewPBFTNetwork(size int) *Blockchain {
    blockchain := NewBlockchain()              // Create a new blockchain instance with the genesis block.
    nodes := make([]Node, size)                // Create an array of nodes.
    for i := 0; i < size; i++ {
        nodes[i] = *NewNode(i, i == 0, blockchain) // Initialize each node; the first node is set as primary.
    }
    blockchain.Nodes = nodes                   // Assign nodes to the blockchain.
    return blockchain
}

// Footer: Security Considerations and Architectural Decisions
//
// This implementation of Practical Byzantine Fault Tolerance (PBFT) demonstrates how nodes in a distributed system
// can reach consensus even when some nodes may act maliciously or exhibit arbitrary behavior. The consensus mechanism
// is designed to tolerate Byzantine faults by requiring agreement from at least 2/3 of the nodes before committing a block.
//
// 1. **Cryptographic Hashing**: Each block's integrity is ensured through the use of SHA-256 hashing. This prevents 
//    unauthorized modifications to blocks, as any change in the block's contents would lead to a different hash.
// 
// 2. **Primary Node Role**: The primary node (or leader) is responsible for proposing new blocks. If the primary fails
//    or acts maliciously, the network must implement a "view change" process to elect a new primary (not implemented here).
//
// 3. **2/3 Majority Consensus**: To tolerate Byzantine faults, a proposed block must receive approvals from at least 
//    2/3 of the nodes before being committed. This ensures that even if some nodes act maliciously, they cannot compromise
//    the consistency of the blockchain.
//
// 4. **Block Verification**: Each node verifies proposed blocks by checking both the previous hash link and recalculating
//    the current block's hash. This two-step verification process ensures both continuity in the chain and data integrity.
//
// This implementation is simplified for educational purposes and demonstrates the core principles of PBFT consensus.
// In a production system, more sophisticated techniques for handling node failures, view changes, and message 
// authentication would be required to maintain resilience and security in a real-world distributed network.