// Package main demonstrates the use of the Delegated Proof of Stake (DPoS) consensus algorithm to manage a simple blockchain.
// DPoS is a consensus mechanism that allows token holders to vote for a small number of trusted delegates who are responsible
// for validating transactions and creating new blocks. This approach is designed to increase scalability and transaction throughput
// while maintaining a decentralized governance model. In this example, delegates are elected through voting, and the elected
// delegates add new blocks to the blockchain.
package main

import (
    "fmt"                            // The fmt package is used for formatted I/O, primarily to print output to the console.
    "consensus-algorithms-edu/algorithms/dpos" // Import the DPoS consensus implementation from the consensus-algorithms-edu module.
)

func main() {
    // Initialize the list of delegates and the voters mapping.
    delegates := []string{"Alice", "Bob", "Charlie"}
    voters := map[string]string{}

    // Create a new blockchain using Delegated Proof of Stake (DPoS) with the provided delegates and an initial set of voters.
    blockchain := dpos.NewBlockchain(delegates, voters)

    // Record votes for delegates by voters.
    blockchain.Vote("Voter1", "Alice")
    blockchain.Vote("Voter2", "Bob")
    blockchain.Vote("Voter3", "Alice")
    blockchain.Vote("Voter4", "Charlie")
    blockchain.Vote("Voter5", "Bob")

    // Count the votes and select delegates based on the number of votes each received.
    blockchain.CountVotes()

    // Add blocks to the blockchain using the elected delegates.
    blockchain.AddBlock("First voting data")
    blockchain.AddBlock("Second voting data")

    // Iterate over each block in the blockchain and print the block's details.
    for _, block := range blockchain.Blocks {
        fmt.Printf("Index: %d\nTimestamp: %s\nData: %s\nPrevious Hash: %s\nHash: %s\nDelegate: %s\n\n", 
            block.Index, block.Timestamp, block.Data, block.PrevHash, block.Hash, block.Delegate)
    }
}

// Footer: Overview and Execution Flow
//
// This example demonstrates how the Delegated Proof of Stake (DPoS) consensus algorithm works to create a decentralized
// blockchain system where elected delegates are responsible for adding new blocks. The process simulates a voting mechanism
// where stakeholders vote for their preferred delegates, and the top-ranked delegates are responsible for block production.
//
// Key Steps:
// 1. **Initialization**: The blockchain is initialized using the `dpos.NewBlockchain()` function, providing a list of delegates.
// 2. **Voting Process**: Voters cast their votes for delegates, recorded using the `Vote()` function. The votes are counted 
//    using `CountVotes()`, which sorts the delegates based on votes received.
// 3. **Block Addition**: New blocks are added to the blockchain by the elected delegates using the `AddBlock()` function.
//    Each block includes information about the delegate who validated and added the block to the chain.
// 4. **Blockchain Display**: Finally, details of each block, such as index, timestamp, data, previous hash, current hash, 
//    and the delegate, are printed to show how the blockchain has evolved.
//
// DPoS is a highly efficient consensus mechanism, making it suitable for environments where scalability and low latency are
// critical. This implementation illustrates how decentralized governance can be achieved by delegating responsibilities to
// trusted representatives, thereby balancing efficiency with democratic control of the blockchain network.
