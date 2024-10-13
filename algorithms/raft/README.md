# Raft Consensus Algorithm

Raft is a consensus algorithm designed to be easy to understand while achieving the same goals as Paxos. Raft is often used in distributed systems for managing replicated logs, ensuring that all participating nodes in the network agree on the same series of actions or events. The primary idea behind Raft is to achieve consensus in a distributed cluster by electing a leader node that is responsible for log replication.

## How Raft Works

1. **Nodes and Roles**:
   - **Leader**: Responsible for managing the replicated log, handling client requests, and distributing them to other nodes.
   - **Followers**: Passively replicate the log entries as dictated by the leader.
   - **Candidate**: A node that can become a leader by initiating an election.
2. **Phases of Raft**:
   - **Leader Election**: If a leader fails or is unreachable, nodes enter a candidate state and initiate an election to select a new leader.
   - **Log Replication**: The leader accepts commands from clients, adds them to its log, and then replicates these log entries to follower nodes.
   - **Commit and Apply**: Once an entry is safely replicated to a majority of nodes, the leader commits the entry and applies it to the state machine. Followers apply entries once they receive confirmation of commitment.
3. **Heartbeats**:
   - Leaders periodically send heartbeat messages to followers to maintain authority and prevent new elections from occurring.

## Features of Raft

- **Leader-Based Consensus**: Raft simplifies consensus by always having a single leader that handles all requests and manages replication.
- **Partition Tolerance**: Raft can continue to make progress as long as a majority of nodes are available.
- **Log Consistency**: All nodes eventually reach consensus on the same sequence of log entries, ensuring consistency in the distributed state machine.

## Structure of This Implementation

In this folder, we provide a Go implementation of the Raft consensus algorithm. The implementation simulates leader election, log replication, and achieving consensus in a distributed network.

### Files

- **`raft.go`**: Contains the Go implementation of the Raft consensus algorithm.

### Key Elements of the Code

- **Blockchain**: Represents the distributed replicated log.
- **Node**: Represents individual nodes that can be a leader, follower, or candidate.
- **Leader Election**: Nodes can transition between follower, candidate, and leader roles as required to maintain a consistent leader in the cluster.

### Code Example

```go
package main

import (
    "fmt"
    "consensus-algorithms-edu/algorithms/raft"
)

func main() {
    blockchain := raft.NewRaftNetwork(5)

    blockchain.Leader.Lead("First log entry")
    blockchain.Leader.Lead("Second log entry")
    blockchain.Leader.Lead("Third log entry")

    for _, block := range blockchain.Blocks {
        fmt.Printf("Index: %d\nTimestamp: %s\nData: %s\nPrevious Hash: %s\nHash: %s\n\n", 
            block.Index, block.Timestamp, block.Data, block.PrevHash, block.Hash)
    }
}
```

### How to Run the Example

1. **Initialize the Network**: Use `NewRaftNetwork()` to create a new Raft network with multiple nodes.
2. **Leader Election**: Initially, one of the nodes is selected as the leader. If the leader fails, other nodes can initiate an election.
3. **Add Log Entries**: Use the `Lead()` function from the leader node to add log entries, which are replicated to other nodes.

### Advantages of Raft

- **Understandable**: Raft was explicitly designed to be easier to understand compared to other consensus algorithms like Paxos, making it suitable for educational purposes and practical implementations.
- **Strong Leadership**: Raft maintains a strong and clear distinction between leader and follower roles, which simplifies log management and consistency.
- **Fault Tolerance**: Raft can tolerate failures of up to half of the nodes (i.e., it requires a majority of nodes to operate).

### Limitations

- **Single Leader Bottleneck**: The leader handles all the requests, which could become a bottleneck under heavy load.
- **Leader Failures**: When the leader fails, the system must perform a new election, which may introduce temporary unavailability until a new leader is elected.

### Use Cases

- **Replicated State Machines**: Raft is often used to implement consistent replicated state machines in distributed systems.
- **Distributed Key-Value Stores**: Systems like etcd and Consul use Raft to achieve consensus for configuration management and service discovery.

### License

This implementation is licensed under the MIT License.
