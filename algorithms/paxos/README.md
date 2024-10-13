# Paxos Consensus Algorithm

Paxos is a consensus algorithm designed to achieve agreement among distributed systems or nodes even in the presence of faults. It was developed to solve the problem of achieving consensus within a group of unreliable or failing participants. Paxos is commonly used in distributed databases, and it ensures safety and liveness, meaning the network will continue to make progress even in adverse conditions.

## How Paxos Works

1. **Proposers, Acceptors, and Learners**:
   - **Proposers** propose values.
   - **Acceptors** vote on the proposed values and determine which value should be accepted.
   - **Learners** learn the agreed value once consensus is reached.
2. **Phases of Paxos**:
   - **Prepare Phase**: A proposer selects a proposal number and sends a prepare request to a quorum of acceptors. Acceptors respond if they have not already responded to a higher-numbered proposal.
   - **Accept Phase**: If the proposer receives enough positive responses, it sends an accept request to a quorum of acceptors with the proposed value. Acceptors will accept the value if they have not already promised to a higher-numbered proposal.
   - **Commit Phase**: Once a proposal is accepted by a majority of acceptors, learners can learn the value.

Paxos is designed to be robust in the presence of failures. It ensures that a single value is agreed upon, which is crucial in distributed environments.

## Features of Paxos

- **Fault Tolerance**: Paxos can reach consensus even if some of the nodes fail, as long as a majority of nodes are still functioning.
- **Guaranteed Safety**: Paxos guarantees that no two nodes will accept different values, ensuring consistency across the system.
- **Progress**: The system can always make progress as long as a quorum of nodes is available.

## Structure of This Implementation

In this folder, we have implemented a simplified version of Paxos in Go. The key components of this implementation include:

### Files

- **`paxos.go`**: Contains the Go implementation of the Paxos consensus algorithm.

### Key Elements of the Code

- **Node**: Represents a node in the distributed system. Nodes can act as proposers, acceptors, or learners.
- **Proposal**: Represents a proposal initiated by a proposer. A proposal includes a proposal ID, data, and an indication of whether it has been accepted.
- **Blockchain**: Represents the agreed chain of data after reaching consensus.

### Code Example

```go
package main

import (
    "fmt"
    "consensus-algorithms-edu/algorithms/paxos"
)

func main() {
    networkSize := 5
    blockchain := paxos.NewPaxosNetwork(networkSize)

    blockchain.RunPaxos("First distributed system data", 1)
    blockchain.RunPaxos("Second distributed system data", 2)
    blockchain.RunPaxos("Third distributed system data", 3)

    for _, block := range blockchain.Blocks {
        fmt.Printf("Index: %d\nTimestamp: %s\nData: %s\nPrevious Hash: %s\nHash: %s\n\n", 
            block.Index, block.Timestamp, block.Data, block.PrevHash, block.Hash)
    }
}
```

### How to Run the Example

1. **Initialize the Network**: Use `NewPaxosNetwork()` to create a distributed network of nodes that will participate in consensus.
2. **Propose Values**: Use `RunPaxos()` to propose a new value to be agreed upon by the nodes.
3. **Reach Consensus**: The network uses Paxos to reach consensus, and the agreed value is added to the blockchain.

### Advantages of Paxos

- **Consistency**: Paxos ensures consistency among nodes, meaning all nodes in the network eventually agree on the same value.
- **Fault Tolerance**: Paxos can tolerate failures of some nodes and still make progress.
- **High Reliability**: Since Paxos requires a majority of nodes to agree, it ensures that decisions are reliable and free from arbitrary failures.

### Limitations

- **Complexity**: Paxos is conceptually difficult to understand and implement. This implementation simplifies some aspects for educational purposes.
- **Performance**: Paxos may suffer from performance bottlenecks in large-scale systems due to the need for multiple communication rounds between nodes.

### License

This implementation is licensed under the MIT License.
