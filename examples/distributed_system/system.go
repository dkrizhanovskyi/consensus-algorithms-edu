// Package main demonstrates the use of the Paxos consensus algorithm to manage a distributed blockchain.
// Paxos is a consensus algorithm used in distributed systems to achieve agreement among a group of nodes, even in the face of failures.
// This simulation demonstrates how Paxos helps achieve consistent state across nodes, using a blockchain as an example.
// The network consists of multiple nodes that participate in the consensus process to propose and commit new blocks to the blockchain.
package main

import (
    "fmt"                             // The fmt package is used for formatted I/O, primarily to print output to the console.
    "consensus-algorithms-edu/algorithms/paxos" // Import the Paxos consensus implementation from the consensus-algorithms-edu module.
)

func main() {
    // Initialize a Paxos network with 5 nodes.
    networkSize := 5
    blockchain := paxos.NewPaxosNetwork(networkSize)

    // Run the Paxos consensus to add data to the blockchain.
    blockchain.RunPaxos("First distributed system data", 1)
    blockchain.RunPaxos("Second distributed system data", 2)
    blockchain.RunPaxos("Third distributed system data", 3)

    // Iterate over each block in the blockchain and print the block's details.
    for _, block := range blockchain.Blocks {
        fmt.Printf("Index: %d\nTimestamp: %s\nData: %s\nPrevious Hash: %s\nHash: %s\n\n", 
            block.Index, block.Timestamp, block.Data, block.PrevHash, block.Hash)
    }
}

// Footer: Overview and Execution Flow
//
// This example demonstrates the use of the Paxos consensus algorithm to create a consistent blockchain across multiple nodes.
// Paxos is well-suited for environments where nodes may fail or where multiple nodes need to agree on a value.
// The primary purpose of this example is to illustrate how distributed consensus is achieved and how a consistent blockchain state is maintained.
//
// Key Steps:
// 1. **Network Initialization**: The Paxos network is initialized with 5 nodes using `paxos.NewPaxosNetwork(networkSize)`.
// 2. **Paxos Proposal and Consensus**: The Paxos consensus algorithm is used to propose and reach agreement on three sets of data.
//    Each proposal results in a new block being created and appended to the blockchain if a majority of nodes agree.
// 3. **Consensus and Fault Tolerance**: During the Paxos process, proposals are broadcast to all nodes. A majority of nodes must approve a proposal before it is committed to the blockchain.
//    This ensures fault tolerance and maintains the integrity of the distributed system.
// 4. **Blockchain Display**: The details of each block, including the index, timestamp, data, previous hash, and current hash, are printed to verify the state of the blockchain.
//
// Paxos is effective for distributed consensus, especially when the network involves multiple nodes with potential failures or delays.
// This implementation highlights the process of achieving consensus in a distributed system, demonstrating the resilience and consistency of Paxos.
