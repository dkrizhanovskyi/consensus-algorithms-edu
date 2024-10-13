# Practical Byzantine Fault Tolerance (PBFT)

Practical Byzantine Fault Tolerance (PBFT) is a consensus algorithm designed to work in distributed systems with the potential for Byzantine faults, which refer to arbitrary or malicious failures of nodes. PBFT provides a way to reach consensus even in the presence of nodes that may act in unexpected, arbitrary, or even intentionally malicious ways. PBFT is particularly well-suited for environments like permissioned blockchains, where a known set of participants is expected to maintain the integrity of the system.

## How PBFT Works

PBFT operates through a series of phases involving a **primary node** (sometimes called a leader) and multiple **replica nodes**. The algorithm proceeds in rounds, with each round having the goal of getting agreement from all honest nodes on a proposed value.

### Roles in PBFT

1. **Primary Node (Leader)**: The node that initiates the proposal. This role can be rotated among all nodes.
2. **Replica Nodes**: Nodes that verify the proposal and participate in reaching consensus.

### Phases of PBFT

1. **Pre-Prepare Phase**:
   - The primary node proposes a new block and sends a pre-prepare message to all replica nodes.
2. **Prepare Phase**:
   - Each replica verifies the pre-prepare message from the primary. If it finds the proposal valid, it broadcasts a prepare message to all other nodes.
   - The prepare phase ensures that all nodes have received the same proposed value from the primary.
3. **Commit Phase**:
   - After receiving a majority of prepare messages, nodes broadcast a commit message to all other nodes.
   - Once a node receives enough commit messages, it commits the proposed value to the blockchain, guaranteeing that all honest nodes reach consensus on the same value.

The protocol ensures that consensus is reached if at least **2f + 1** nodes (where **f** is the maximum number of faulty nodes that can be tolerated) agree on the value. This implies that PBFT can tolerate **f** Byzantine failures in a network of **3f + 1** nodes.

## Key Features of PBFT

- **Byzantine Fault Tolerance**: PBFT can tolerate nodes that exhibit arbitrary (Byzantine) behavior, making it well-suited for environments with a risk of internal faults or malicious actors.
- **High Throughput**: PBFT is designed for high throughput by eliminating the need for resource-intensive mining processes used in PoW.
- **Deterministic Finality**: Once consensus is reached, the decision is final. There is no risk of forks or chain reorganizations, unlike in PoW-based blockchains.

## Structure of This Implementation

In this folder, we have implemented a simplified version of PBFT in Go. The implementation allows you to see how a group of nodes can reach consensus in the presence of potential Byzantine faults through message exchange and verification.

### Files

- **`pbft.go`**: Contains the Go implementation of the PBFT consensus algorithm.

### Key Elements of the Code

- **Blockchain**: Represents the chain of blocks that the nodes reach consensus on.
- **Primary and Replica Nodes**: Nodes play roles as either primary or replicas. The primary node proposes values, and replica nodes participate in verification and consensus.
- **Phases of PBFT**: The code demonstrates the pre-prepare, prepare, and commit phases used to achieve consensus.

### Code Example

```go
package main

import (
    "fmt"
    "consensus-algorithms-edu/algorithms/pbft"
)

func main() {
    blockchain := pbft.NewPBFTNetwork(5)

    blockchain.RunPBFT("First transaction data")
    blockchain.RunPBFT("Second transaction data")
    blockchain.RunPBFT("Third transaction data")

    for _, block := range blockchain.Blocks {
        fmt.Printf("Index: %d\nTimestamp: %s\nData: %s\nPrevious Hash: %s\nHash: %s\n\n", 
            block.Index, block.Timestamp, block.Data, block.PrevHash, block.Hash)
    }
}
```

### How to Run the Example

1. **Initialize the PBFT Network**: Use `NewPBFTNetwork()` to create a new PBFT network with multiple nodes.
2. **Run PBFT Consensus**: Use `RunPBFT()` to propose a value, which is then verified and committed by the network using PBFT phases.
3. **View Consensus Results**: The committed blocks are then printed to show the successful consensus.

### Advantages of PBFT

- **Byzantine Fault Tolerance**: PBFT is resilient against nodes that may behave arbitrarily or maliciously, as long as a majority of nodes are honest.
- **Energy Efficiency**: Unlike PoW, PBFT does not require extensive computational power, making it suitable for energy-constrained environments.
- **Finality**: Once consensus is reached, there is immediate and deterministic finality, which means that the block cannot be changed or reorganized.

### Limitations

- **Communication Complexity**: PBFT requires multiple rounds of message exchanges among nodes, which means that it has high communication overhead, especially as the number of nodes grows.
- **Scalability**: PBFT is less suitable for large-scale systems due to its communication requirements. It works best in permissioned systems with a smaller number of nodes.
- **Leader Dependence**: PBFT relies on a primary node (leader), and its performance may degrade if the primary node behaves maliciously or becomes unreliable. In such cases, view changes are needed to elect a new leader.

## Practical Use Cases of PBFT

- **Permissioned Blockchains**: PBFT is used in blockchain platforms that have a fixed set of known participants, such as Hyperledger Fabric.
- **Financial Systems**: PBFT is suitable for use in financial applications where security and trust are paramount, and where participants are vetted.
- **Distributed Databases**: PBFT can be used to manage replicated databases that require high levels of consistency across multiple nodes.

## Comparison with Other Consensus Algorithms

### PBFT vs. Proof of Work (PoW)
- **Efficiency**: PBFT is far more efficient compared to PoW, as it does not involve energy-intensive mining.
- **Latency**: PBFT achieves faster consensus due to its leader-based structure, whereas PoW introduces significant latency due to mining difficulty.
- **Fault Tolerance**: PBFT can handle Byzantine faults, whereas PoW is primarily resilient to Sybil attacks and requires computational proof to deter attackers.

### PBFT vs. Raft
- **Fault Model**: PBFT is designed to tolerate Byzantine faults, whereas Raft assumes fail-stop failures where nodes may crash but not behave maliciously.
- **Complexity**: PBFT is more complex due to its multi-phase communication protocol and its ability to handle Byzantine nodes. Raft is simpler and designed to be easy to implement and understand.

### Technical Overview

The PBFT implementation provided in this project demonstrates the following key aspects:

1. **Leader-Based Proposal**: The primary node proposes a block, which is then validated by replica nodes.
2. **Prepare and Commit Phases**: The prepare phase ensures that all nodes have received the proposal, and the commit phase ensures that the value is finalized and replicated across the network.
3. **Consensus Finality**: Once nodes reach consensus, the value is committed, and the ledger remains consistent across all honest nodes.

### Example Workflow

- **Pre-Prepare Phase**: The primary node proposes a new block containing transaction data.
- **Prepare Phase**: All replica nodes verify the proposal and broadcast their acceptance to others.
- **Commit Phase**: After receiving enough prepare messages, nodes broadcast commit messages. Once the commit threshold is reached, the value is added to the blockchain.

### License

This documentation and the associated code are licensed under the MIT License.
