# Practical Byzantine Fault Tolerance (PBFT)

Practical Byzantine Fault Tolerance (PBFT) is a consensus algorithm designed to work in distributed systems with potential Byzantine faults, which refer to arbitrary or malicious failures. PBFT is used to reach consensus even when some nodes in the network act maliciously or provide incorrect information. PBFT is particularly well-suited for scenarios where nodes may be compromised, such as in blockchain networks.

## How PBFT Works

1. **Network Participants**:
   - **Primary Node**: Also known as the leader, initiates proposals.
   - **Replica Nodes**: Other nodes that verify and agree on the correctness of the proposals.
2. **Phases of PBFT**:
   - **Pre-Prepare Phase**: The primary node proposes a request by broadcasting it to all replica nodes.
   - **Prepare Phase**: Each replica verifies the request and broadcasts a prepare message to the network.
   - **Commit Phase**: Each node broadcasts a commit message after receiving enough prepare messages. If a replica receives enough commit messages, it commits the request to the blockchain.
3. **Consensus**:
   - To reach consensus, at least `2f + 1` nodes (where `f` is the number of faulty nodes tolerated) must agree on the same value.
   - PBFT is effective in systems with a small number of malicious nodes.

## Features of PBFT

- **Byzantine Fault Tolerance**: PBFT can tolerate malicious or faulty nodes, making it ideal for use in environments where some nodes may not behave correctly.
- **Low Latency**: Compared to Proof of Work (PoW), PBFT has lower latency since it does not require extensive computational resources to solve complex puzzles.
- **Deterministic Finality**: Once consensus is reached, the value is immediately final and cannot be reverted.

## Structure of This Implementation

In this folder, we provide a Go implementation of the PBFT consensus algorithm. The implementation demonstrates how nodes work together to reach consensus even in the presence of Byzantine faults.

### Files

- **`pbft.go`**: Contains the Go implementation of the Practical Byzantine Fault Tolerance consensus algorithm.

### Key Elements of the Code

- **Blockchain**: Represents the chain of blocks agreed upon by the nodes.
- **Node**: Represents individual nodes that participate in consensus. Nodes can be primary or replica nodes.
- **Phases**: The implementation simulates the Pre-Prepare, Prepare, and Commit phases to reach consensus.

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

1. **Initialize the Network**: Use `NewPBFTNetwork()` to create a distributed network of nodes.
2. **Propose Values**: Use `RunPBFT()` to propose a new value that all nodes must reach consensus on.
3. **Phases of PBFT**: The algorithm will go through Pre-Prepare, Prepare, and Commit phases to reach consensus.

### Advantages of PBFT

- **Byzantine Fault Tolerance**: PBFT can tolerate up to `f` Byzantine faults in a network with `3f + 1` nodes, making it robust against malicious actors.
- **Efficiency**: Unlike PoW, PBFT does not require computational resources for mining, leading to faster consensus.
- **Finality**: Once consensus is achieved, the value is final, and there is no risk of chain reorganization.

### Limitations

- **Scalability**: PBFT works well in systems with a relatively small number of nodes. As the number of nodes increases, the communication overhead also increases.
- **Complex Communication**: The consensus requires multiple rounds of communication between nodes, which can be complex and lead to delays if there are too many nodes.

### License

This implementation is licensed under the MIT License.
