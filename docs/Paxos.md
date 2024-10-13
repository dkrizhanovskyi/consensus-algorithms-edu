# Paxos Consensus Algorithm

Paxos is a consensus algorithm designed to enable a group of distributed nodes to reach agreement on a single value, even in the presence of failures. It was developed by Leslie Lamport and is one of the foundational algorithms used in distributed systems to achieve fault-tolerant consensus. Paxos is designed to ensure safety (no two nodes decide on different values) and liveness (the network eventually reaches consensus).

## How Paxos Works

Paxos operates through three main roles: **Proposers**, **Acceptors**, and **Learners**. The algorithm consists of several phases to propose, accept, and commit values.

### Roles in Paxos

1. **Proposers**: Nodes that propose values for consensus.
2. **Acceptors**: Nodes that vote on proposed values and are responsible for accepting proposals.
3. **Learners**: Nodes that learn the agreed value after consensus is reached.

### Phases of Paxos

1. **Prepare Phase**:
   - A proposer selects a unique proposal number and sends a "prepare" request to a quorum of acceptors.
   - If an acceptor has not already responded to a proposal with a higher number, it responds with a "promise" not to accept lower-numbered proposals and provides any previously accepted proposal value.

2. **Accept Phase**:
   - If the proposer receives promises from a majority of acceptors, it sends an "accept" request with a value.
   - Acceptors then accept the proposal unless they have already promised to accept a higher-numbered proposal.

3. **Commit Phase**:
   - Once a value is accepted by a majority of acceptors, the value is committed, and learners are notified of the final decision.

Paxos ensures that consensus is achieved even in the presence of a minority of faulty nodes by requiring that a majority of nodes agree on each phase of the protocol.

## Key Features of Paxos

- **Fault Tolerance**: Paxos is designed to tolerate failures in up to half of the nodes in the network, as long as a majority are still functioning correctly.
- **Consensus with Multiple Proposers**: Paxos allows multiple nodes to propose values concurrently, and the protocol guarantees that only one value is chosen.
- **Guaranteed Safety**: Paxos ensures that no two nodes will decide on different values, even if network partitions or failures occur.

## Structure of This Implementation

In this folder, we provide a simplified version of Paxos implemented in Go. The implementation demonstrates the basic process of proposing values, achieving agreement among acceptors, and committing the decided value.

### Files

- **`paxos.go`**: Contains the Go implementation of the Paxos consensus algorithm.

### Key Elements of the Code

- **Proposers, Acceptors, Learners**: The code simulates the roles of proposers, acceptors, and learners, enabling them to communicate and reach agreement.
- **Proposal Process**: A proposal consists of a unique ID and a value. Proposers submit proposals, and acceptors decide whether to accept them based on the proposal ID.
- **Commitment**: Once a proposal is accepted by a majority of acceptors, it is committed and all learners are informed.

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

1. **Initialize the Network**: Use `NewPaxosNetwork()` to create a network of nodes capable of running the Paxos protocol.
2. **Propose Values**: Use the `RunPaxos()` method to propose a new value that needs to be agreed upon by the network.
3. **Observe Consensus**: The value proposed will go through the prepare, accept, and commit phases, and eventually be added to the blockchain.

### Advantages of Paxos

- **Robustness**: Paxos is resilient to node failures and ensures consistency even when multiple nodes propose conflicting values.
- **Safety Guarantee**: Paxos ensures that at most one value is chosen, which guarantees that all nodes will eventually agree on a single value.
- **Fault Tolerance**: As long as a majority of nodes are available, Paxos can continue to make progress.

### Limitations

- **Complexity**: Paxos is known for its complexity, both conceptually and in terms of implementation. This complexity can make it difficult to understand and prone to bugs.
- **Performance**: Paxos involves multiple rounds of communication, which can introduce significant latency in large-scale distributed systems.
- **Coordination Overhead**: Since Paxos requires a majority agreement, the coordination among nodes can lead to delays, especially in geographically distributed networks.

## Practical Use Cases of Paxos

- **Distributed Databases**: Paxos is widely used in distributed databases and systems like Google Spanner and Amazon Dynamo to ensure consistency among replicas.
- **Leader Election**: Paxos is often used as a protocol for leader election in distributed systems to ensure that a single leader is chosen in a fault-tolerant manner.
- **Fault-Tolerant Services**: Paxos is employed in building services that must continue to operate correctly even when some components fail, such as consensus among replicas in cloud services.

## Comparison with Other Consensus Algorithms

### Paxos vs. Raft
- **Understandability**: Raft was designed to be easier to understand and implement compared to Paxos, which is why Raft has gained popularity in recent years.
- **Leader Election**: Both algorithms use leader election as part of their approach, but Raft maintains a single leader at all times, while Paxos can handle multiple proposers.
- **Complexity**: Paxos is conceptually more complex and has been historically challenging to implement correctly, whereas Raft explicitly aims to simplify consensus.

### Paxos vs. Byzantine Fault Tolerance (BFT)
- **Fault Tolerance**: Paxos is designed to handle fail-stop failures, while Byzantine Fault Tolerance algorithms, like PBFT, can handle arbitrary or malicious faults.
- **Consensus Guarantees**: Paxos guarantees consensus among nodes that may crash or become unresponsive, whereas BFT algorithms are more resilient to malicious actors.

## Technical Overview

The Paxos implementation in this repository is simplified to help understand its core components:

- **Prepare, Accept, and Commit Phases**: The phases are implemented in a straightforward way to demonstrate how nodes interact to achieve consensus.
- **Majority Agreement**: Paxos requires a majority of nodes to agree in each phase to move forward, ensuring that the system can always make progress despite failures.

### Example Workflow

- **Proposer Node** sends a "prepare" request to acceptors.
- **Acceptors** respond to the request, and if a majority agree, the proposer sends an "accept" request with the value.
- Once the value is **accepted** by a majority, the **learners** are informed, and the value is committed.

### License

This documentation and the associated code are licensed under the MIT License.
