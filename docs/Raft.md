# Raft Consensus Algorithm

Raft is a consensus algorithm designed to be easier to understand than other consensus algorithms like Paxos, while maintaining similar functionality and reliability. Raft is often used to manage replicated logs in distributed systems, ensuring that all nodes in a cluster agree on the state and order of commands, even in the face of network failures and node crashes. Raft achieves this by electing a leader node and using that leader to coordinate log replication across follower nodes.

## How Raft Works

Raft organizes the consensus process into multiple distinct roles and phases to achieve fault-tolerant consensus. Each node in a Raft cluster can be in one of three states: **Leader**, **Follower**, or **Candidate**.

### Roles in Raft

1. **Leader**: The node responsible for managing the replicated log and coordinating with followers. It handles client requests and distributes log entries to the follower nodes.
2. **Follower**: Passive nodes that replicate the log entries as directed by the leader.
3. **Candidate**: A follower can become a candidate when it times out without hearing from the leader and initiates an election to become the new leader.

### Phases of Raft

1. **Leader Election**:
   - If a follower does not receive communication from the leader within a certain timeout period, it assumes the leader has failed and becomes a candidate.
   - The candidate initiates an election by requesting votes from other nodes. Each node can vote for at most one candidate per term.
   - If the candidate receives a majority of the votes, it becomes the new leader. Otherwise, the process repeats until a leader is elected.

2. **Log Replication**:
   - Once a leader is established, it handles all client requests and appends them to its log.
   - The leader then sends **AppendEntries** messages to follower nodes to replicate the log entries. Followers acknowledge receipt, and once a majority of nodes have confirmed the entry, the leader commits it to the log.

3. **Commit and Apply**:
   - After a log entry is replicated to a majority of nodes, the leader marks it as committed and applies it to the state machine.
   - Followers also apply committed entries to their state machine to ensure consistency across the cluster.

4. **Heartbeats**:
   - The leader regularly sends heartbeat messages to followers to prevent them from becoming candidates. Heartbeats ensure the followers know that the leader is still active.

## Key Features of Raft

- **Leader-Based Consensus**: At any given time, there is a single leader responsible for managing all the operations, which simplifies the consensus process.
- **Fault Tolerance**: Raft is designed to tolerate failures of nodes, allowing the system to continue operating as long as a majority of nodes are available.
- **Log Consistency**: Raft ensures that all nodes have the same sequence of log entries, even in the face of network partitions or node crashes.

## Structure of This Implementation

In this folder, we have provided a simplified version of the Raft consensus algorithm implemented in Go. This implementation demonstrates leader election, log replication, and consensus achievement in a distributed system.

### Files

- **`raft.go`**: Contains the Go implementation of the Raft consensus algorithm.

### Key Elements of the Code

- **Node**: Represents an individual node that can be a follower, leader, or candidate.
- **Leader Election**: Implements the process by which a node becomes the leader.
- **Log Replication**: The leader node is responsible for replicating log entries to all follower nodes.

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

1. **Initialize the Raft Network**: Use `NewRaftNetwork()` to create a Raft network with multiple nodes.
2. **Leader Election**: Initially, one of the nodes is chosen as the leader. If this leader fails, another node is elected as the leader.
3. **Add Log Entries**: Use the `Lead()` function to add log entries, which are then replicated across the cluster.

### Advantages of Raft

- **Simplicity**: Raft was explicitly designed to be easier to understand compared to Paxos, making it suitable for real-world implementations and educational purposes.
- **Strong Leadership**: Raft simplifies the consensus process by establishing a clear leader, which eliminates conflicts and simplifies decision-making.
- **Fault Tolerance**: Raft can tolerate up to half of the nodes failing while continuing to operate, providing robust fault tolerance in distributed environments.

### Limitations

- **Single Leader Bottleneck**: Since Raft relies on a single leader to handle all client requests, the leader can become a bottleneck if there is a high volume of requests.
- **Election Overhead**: In the event of a leader failure, Raft must conduct an election, which may introduce temporary downtime until a new leader is elected.

## Practical Use Cases of Raft

- **Distributed Databases**: Raft is commonly used in distributed databases to ensure consistency among replicas. It powers consensus in systems like etcd, Consul, and RethinkDB.
- **Cluster Management**: Raft is also used in cluster management tools to maintain consistent state and facilitate leader election for cluster operations.
- **Log Replication**: Raft provides a straightforward way to manage replicated logs, ensuring that all nodes in the system have the same view of the data.

## Comparison with Other Consensus Algorithms

### Raft vs. Paxos
- **Understandability**: Raft was designed to be easier to understand and implement compared to Paxos, which is more abstract and complex. Raft explicitly uses the concept of leadership to simplify consensus.
- **Leadership Model**: Raft always maintains a single, strong leader responsible for coordinating all actions, whereas Paxos can involve multiple proposers, which adds complexity.

### Raft vs. Practical Byzantine Fault Tolerance (PBFT)
- **Fault Tolerance**: Raft is designed to tolerate crash failures but does not handle Byzantine faults (nodes behaving arbitrarily or maliciously). PBFT is designed to handle Byzantine faults, which makes it more resilient to malicious behavior but also more complex.
- **Scalability**: Raft is simpler and more suitable for smaller distributed clusters. PBFT involves significantly more communication between nodes, which can be less efficient for large networks.

### Technical Overview

The Raft implementation in this project demonstrates the following:

1. **Leader Election**: Nodes initiate elections when they lose contact with the leader. The election process ensures that a new leader is chosen and can continue managing the cluster.
2. **Log Replication**: The leader is responsible for replicating log entries to all follower nodes, and these entries are committed once a majority of followers acknowledge them.
3. **Heartbeat Mechanism**: Leaders send regular heartbeat messages to keep followers from timing out and starting new elections.

### Example Workflow

- **Leader Election**: When the current leader is unresponsive, the nodes in the network initiate an election to select a new leader.
- **Client Requests**: The leader handles incoming client requests, appends the commands to its log, and replicates them to follower nodes.
- **Log Commitment**: Once a majority of nodes have appended the new entry, it is marked as committed, and the leader and followers apply the changes to their state machines.

### License

This documentation and the associated code are licensed under the MIT License.
