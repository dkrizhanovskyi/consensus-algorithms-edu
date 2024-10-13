// Package raft implements a basic version of the Raft consensus algorithm for a distributed blockchain network.
// Raft is a consensus algorithm designed for managing replicated logs in distributed systems, focusing on simplicity
// and understandability compared to other consensus algorithms like Paxos. It relies on leader election to ensure
// that only one node is actively managing updates, thereby simplifying the consensus process.
// This implementation simulates the core functions of Raft, including leader election, proposing new blocks, and 
// achieving consensus before committing blocks to the blockchain.
package raft

import (
    "crypto/sha256"
    "fmt"
    "strconv"
    "time"
)

// Block represents an individual block in the blockchain.
// It contains information such as the index, timestamp, data, and cryptographic hashes.
type Block struct {
    Index     int    // Position of the block in the blockchain.
    Timestamp string // Time when the block was created.
    Data      string // Data contained within the block (e.g., transactions).
    PrevHash  string // Hash of the previous block to maintain immutability.
    Hash      string // SHA-256 hash of the current block.
}

// Blockchain represents the distributed ledger that is managed by multiple nodes.
type Blockchain struct {
    Blocks []Block // A slice of all blocks in the blockchain.
    Nodes  []Node  // A list of nodes participating in the Raft consensus network.
    Leader *Node   // Pointer to the current leader node responsible for managing updates.
}

// Node represents an individual node within the Raft network.
// Nodes can participate in leader elections, propose blocks, and verify or commit blocks.
type Node struct {
    ID         int         // Unique identifier for the node.
    IsLeader   bool        // Indicates if the node is the leader.
    Blockchain *Blockchain // Reference to the blockchain managed by the node.
}

// NewBlock creates a new block given data, the previous block's hash, and the index.
// It calculates the block's hash to ensure integrity.
func NewBlock(data string, prevHash string, index int) Block {
    block := Block{
        Index:     index,
        Timestamp: time.Now().String(), // Set the current timestamp for the block.
        Data:      data,
        PrevHash:  prevHash,
    }
    block.Hash = block.CalculateHash() // Calculate the cryptographic hash for the new block.
    return block
}

// CalculateHash generates the SHA-256 hash of the block's contents.
// This ensures that any change to the block's data will produce a completely different hash.
func (b *Block) CalculateHash() string {
    record := strconv.Itoa(b.Index) + b.Timestamp + b.Data + b.PrevHash
    hash := sha256.New()                // Create a new SHA-256 hash object.
    hash.Write([]byte(record))          // Write the concatenated block data to the hash.
    hashed := hash.Sum(nil)             // Compute the hash value.
    return fmt.Sprintf("%x", hashed)    // Return the hash as a hexadecimal string.
}

// AddBlock appends a new block to the blockchain.
// This function is called once a new block is validated and consensus is achieved.
func (bc *Blockchain) AddBlock(block Block) {
    bc.Blocks = append(bc.Blocks, block) // Append the new block to the blockchain.
}

// NewBlockchain initializes a new blockchain with a genesis block.
// The genesis block is the initial block that forms the foundation of the blockchain.
func NewBlockchain() *Blockchain {
    genesisBlock := NewBlock("Genesis Block", "", 0) // Create the genesis block (index 0).
    return &Blockchain{
        Blocks: []Block{genesisBlock}, // Initialize with the genesis block.
        Nodes:  []Node{},              // Initialize an empty list of nodes.
    }
}

// ProposeBlock allows the leader node to create a new block proposal based on the latest block.
func (n *Node) ProposeBlock(data string) Block {
    prevBlock := n.Blockchain.Blocks[len(n.Blockchain.Blocks)-1] // Retrieve the latest block.
    newBlock := NewBlock(data, prevBlock.Hash, prevBlock.Index+1) // Create a new block with the provided data.
    return newBlock
}

// BroadcastBlock sends a proposed block to all nodes for verification.
// A block is considered valid if more than half of the nodes approve it.
func (bc *Blockchain) BroadcastBlock(block Block) bool {
    approvals := 0
    totalNodes := len(bc.Nodes)
    
    for _, node := range bc.Nodes {
        if node.VerifyBlock(block) {
            approvals++ // Count nodes that approve the block.
        }
    }
    
    return approvals > totalNodes/2 // Return true if a majority of nodes approve the block.
}

// VerifyBlock allows a node to verify the validity of a proposed block.
// It checks if the previous hash matches the last block in the chain and if the block hash is correct.
func (n *Node) VerifyBlock(block Block) bool {
    prevBlock := n.Blockchain.Blocks[len(n.Blockchain.Blocks)-1] // Retrieve the latest block.
    // Check if the proposed block's previous hash matches the latest block and if the hash is valid.
    if block.PrevHash == prevBlock.Hash {
        return block.Hash == block.CalculateHash()
    }
    return false
}

// CommitBlock commits a verified block to the blockchain.
// This function is called by all nodes once consensus has been achieved.
func (n *Node) CommitBlock(block Block) {
    n.Blockchain.AddBlock(block) // Append the verified block to the blockchain.
}

// RequestVote allows a node to request votes from other nodes during the leader election process.
// If the node receives a majority of votes, it becomes the new leader.
func (n *Node) RequestVote() bool {
    votes := 0
    totalNodes := len(n.Blockchain.Nodes)
    
    for _, node := range n.Blockchain.Nodes {
        if node.VoteFor(n.ID) {
            votes++ // Count votes received from other nodes.
        }
    }
    
    if votes > totalNodes/2 {
        n.IsLeader = true            // Node becomes the leader if it receives a majority of votes.
        n.Blockchain.Leader = n      // Update the blockchain's leader reference.
        return true
    }
    return false
}

// VoteFor allows a node to vote for a candidate during the leader election.
// In this simplified version, nodes always vote for the requesting candidate.
func (n *Node) VoteFor(candidateID int) bool {
    return true // Simplified: Always vote in favor of the candidate.
}

// Lead allows the leader to propose and commit a new block to the blockchain.
// The leader proposes a block, broadcasts it for approval, and if approved, commits it.
func (n *Node) Lead(data string) {
    if n.IsLeader {
        newBlock := n.ProposeBlock(data) // Leader proposes a new block.
        // Broadcast the proposed block and commit it if approved by the majority.
        if n.Blockchain.BroadcastBlock(newBlock) {
            for _, node := range n.Blockchain.Nodes {
                node.CommitBlock(newBlock)
            }
        }
    }
}

// NewNode creates a new node with the given ID and associates it with a blockchain.
func NewNode(id int, blockchain *Blockchain) *Node {
    return &Node{
        ID:         id,
        IsLeader:   false,       // By default, nodes are not leaders initially.
        Blockchain: blockchain,
    }
}

// NewRaftNetwork initializes a Raft network with the specified number of nodes.
// The nodes collaborate to reach consensus and elect a leader to manage block proposals.
func NewRaftNetwork(size int) *Blockchain {
    blockchain := NewBlockchain()              // Create a new blockchain instance.
    nodes := make([]Node, size)                // Create an array of nodes.
    for i := 0; i < size; i++ {
        nodes[i] = *NewNode(i, blockchain)     // Initialize each node and link it to the blockchain.
    }
    blockchain.Nodes = nodes                   // Assign the nodes to the blockchain.
    return blockchain
}

// Footer: Security Considerations and Architectural Decisions
//
// This implementation of Raft demonstrates the basic principles of leader election, consensus, and block management.
// Raft is a consensus algorithm that focuses on simplicity and partition tolerance, making it easier to implement and understand.
//
// 1. **Leader Election**: Nodes can request votes to become the leader. The leader is responsible for managing all updates
//    and ensuring that only consistent blocks are added to the blockchain. Leader election ensures that there is only one
//    authoritative node managing updates, reducing conflicts.
//
// 2. **Majority Consensus**: For a block to be added to the blockchain, a majority of nodes must verify and approve it.
//    This majority consensus ensures the integrity of the blockchain, even if some nodes are compromised or offline.
//
// 3. **Simplified Voting**: In this implementation, nodes always vote for the requesting candidate during leader election.
//    In a real-world scenario, additional rules such as candidate term limits and persistent leader information would
//    prevent conflicts and ensure stability.
//
// 4. **Data Integrity**: Each block's hash is computed based on the previous hash, timestamp, data, and other block metadata.
//    This hash linkage ensures immutability and consistency, as altering any data would require recalculating all subsequent blocks.
//
// Raft is a robust consensus mechanism that provides fault tolerance, making it suitable for distributed systems like databases and
// cluster management tools. This implementation is a simplified educational version to help understand the key concepts
// behind Raft's leader-based consensus model.