// Package paxos implements a basic version of the Paxos consensus algorithm, which is widely used in distributed systems
// to achieve consensus among multiple nodes, even in the presence of failures. Paxos is particularly effective in ensuring
// consistency in distributed databases and replicated state machines, where achieving agreement on a single value is critical
// for maintaining system reliability. This module simulates a blockchain network where nodes propose and commit blocks
// using a simplified version of the Paxos protocol.
package paxos

import (
    "crypto/sha256"
    "fmt"
    "strconv"
    "time"
)

// Block represents an individual block in the blockchain.
// Each block includes metadata, cryptographic hashes, and the data it contains.
type Block struct {
    Index     int    // Position of the block in the blockchain.
    Timestamp string // Timestamp of when the block was created.
    Data      string // Data held within the block, typically transaction details.
    PrevHash  string // The hash of the previous block, ensuring continuity of the chain.
    Hash      string // The cryptographic hash of the current block.
}

// Blockchain represents the distributed ledger managed by nodes participating in the Paxos consensus process.
type Blockchain struct {
    Blocks []Block // Slice containing all the blocks in the blockchain.
    Nodes  []Node  // Slice representing all nodes participating in the Paxos consensus.
}

// Node represents a participant in the Paxos network.
// Each node can propose values and participate in accepting proposals to reach consensus.
type Node struct {
    ID         int         // Unique identifier for the node.
    Proposals  []Proposal  // List of proposals that this node has made or accepted.
    Blockchain *Blockchain // Reference to the blockchain managed by this node.
}

// Proposal represents a proposed value that nodes can either accept or reject.
type Proposal struct {
    ProposalID int    // Unique identifier for the proposal.
    Data       string // The data being proposed for consensus.
    Accepted   bool   // Flag indicating if the proposal has been accepted.
}

// NewBlock creates a new block with the provided data, index, and reference to the previous block's hash.
// The new block's hash is calculated to ensure integrity.
func NewBlock(data string, prevHash string, index int) Block {
    block := Block{
        Index:     index,
        Timestamp: time.Now().String(), // Timestamp to record when the block was created.
        Data:      data,
        PrevHash:  prevHash,
    }
    block.Hash = block.CalculateHash() // Calculate and assign the cryptographic hash for the block.
    return block
}

// CalculateHash generates a cryptographic SHA-256 hash of the block's contents to ensure immutability.
func (b *Block) CalculateHash() string {
    record := strconv.Itoa(b.Index) + b.Timestamp + b.Data + b.PrevHash
    hash := sha256.New()                // Create a new SHA-256 hash object.
    hash.Write([]byte(record))          // Write the concatenated block data to the hash.
    hashed := hash.Sum(nil)             // Compute the hash value.
    return fmt.Sprintf("%x", hashed)    // Return the hash value as a hexadecimal string.
}

// AddBlock appends a new block to the blockchain.
func (bc *Blockchain) AddBlock(block Block) {
    bc.Blocks = append(bc.Blocks, block) // Append the new block to the chain.
}

// NewBlockchain initializes a new blockchain with a genesis block.
// The genesis block serves as the foundation of the chain and is always the first block.
func NewBlockchain() *Blockchain {
    genesisBlock := NewBlock("Genesis Block", "", 0) // Create the genesis block.
    return &Blockchain{
        Blocks: []Block{genesisBlock}, // Initialize with the genesis block.
        Nodes:  []Node{},              // Initialize an empty list of nodes.
    }
}

// Propose allows a node to create a new proposal containing data to be added to the blockchain.
// The proposal is recorded for potential consensus.
func (n *Node) Propose(data string, proposalID int) Proposal {
    proposal := Proposal{
        ProposalID: proposalID,
        Data:       data,
        Accepted:   false,
    }
    n.Proposals = append(n.Proposals, proposal) // Add the new proposal to the list of proposals.
    return proposal
}

// BroadcastProposal broadcasts the given proposal to all nodes in the blockchain network.
// Each node decides whether to accept the proposal. The proposal is accepted if more than half of the nodes agree.
func (bc *Blockchain) BroadcastProposal(proposal Proposal) bool {
    approvals := 0
    totalNodes := len(bc.Nodes)
    
    for _, node := range bc.Nodes {
        if node.AcceptProposal(proposal) {
            approvals++ // Count nodes that accept the proposal.
        }
    }
    
    // Return true if the majority of nodes approve the proposal.
    return approvals > totalNodes/2
}

// AcceptProposal is called by a node to decide if it will accept a given proposal.
// The proposal is marked as accepted and recorded if it matches an existing proposal's ID.
func (n *Node) AcceptProposal(proposal Proposal) bool {
    for _, p := range n.Proposals {
        if p.ProposalID == proposal.ProposalID {
            p.Accepted = true // Mark the proposal as accepted.
            return true
        }
    }
    return false // Return false if the proposal is not found in the node's list of proposals.
}

// CommitProposal commits an accepted proposal to the blockchain.
// This involves creating a new block based on the proposal data and appending it to the chain.
func (n *Node) CommitProposal(proposal Proposal) {
    prevBlock := n.Blockchain.Blocks[len(n.Blockchain.Blocks)-1] // Get the last block in the chain.
    newBlock := NewBlock(proposal.Data, prevBlock.Hash, prevBlock.Index+1)
    n.Blockchain.AddBlock(newBlock)                              // Append the new block to the blockchain.
}

// RunPaxos initiates the Paxos consensus process for the given proposal data and proposal ID.
// The first node in the blockchain proposes the data, and consensus is achieved if a majority approve.
func (bc *Blockchain) RunPaxos(data string, proposalID int) {
    proposer := bc.Nodes[0]                     // Select the first node as the proposer.
    proposal := proposer.Propose(data, proposalID) // Create a new proposal.

    // Broadcast the proposal and, if approved by a majority, commit it to the blockchain.
    if bc.BroadcastProposal(proposal) {
        for _, node := range bc.Nodes {
            node.CommitProposal(proposal)       // Each node commits the approved proposal.
        }
    }
}

// NewNode creates a new node with the given ID and associates it with a blockchain.
func NewNode(id int, blockchain *Blockchain) *Node {
    return &Node{
        ID:         id,
        Blockchain: blockchain,
    }
}

// NewPaxosNetwork initializes a Paxos network with the specified number of nodes.
// Each node is part of the blockchain, and the nodes collaborate to achieve consensus.
func NewPaxosNetwork(size int) *Blockchain {
    blockchain := NewBlockchain()            // Create a new blockchain instance.
    nodes := make([]Node, size)              // Create an array of nodes.
    for i := 0; i < size; i++ {
        nodes[i] = *NewNode(i, blockchain)   // Initialize each node and link it to the blockchain.
    }
    blockchain.Nodes = nodes                 // Assign the nodes to the blockchain.
    return blockchain
}

// Footer: Security Considerations and Architectural Decisions
//
// This implementation of Paxos provides a simplified version of the consensus algorithm for educational purposes,
// simulating how nodes in a distributed system achieve agreement even in the face of potential failures or inconsistencies.
//
// 1. **Cryptographic Hashing**: Each block is hashed using SHA-256 to ensure data integrity and immutability. Modifying
//    any part of a block (such as the index or data) results in a completely different hash, making tampering detectable.
//
// 2. **Paxos Consensus Workflow**: Paxos is divided into distinct steps: proposing values, broadcasting proposals, 
//    and accepting or rejecting those proposals. The broadcasted proposal must be accepted by a majority of nodes 
//    before it is committed, ensuring a robust majority consensus.
//
// 3. **Leader Selection and Simplification**: For simplicity, this implementation always selects the first node as 
//    the proposer. In practical distributed systems, a leader-election mechanism is often employed to dynamically 
//    determine the proposer in case of failures or changes in network conditions.
//
// 4. **Fault Tolerance**: Paxos is designed to tolerate failures by ensuring that proposals are only committed if 
//    a majority of nodes agree. This helps prevent inconsistencies even if some nodes fail or behave erratically