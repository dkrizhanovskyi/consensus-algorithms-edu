# Distributed System Example Using Consensus Algorithms

This folder contains an example of a distributed system that uses different consensus algorithms—such as Paxos and Raft—to maintain consistent state across nodes in the system. This example is designed to provide educational value, demonstrating how distributed consensus works to ensure all nodes agree on a common state, even in the presence of failures.

## Overview

The distributed system example provided here is a simplified version of a distributed key-value store, where nodes participate in reaching consensus to ensure data consistency across the entire system. Consensus is crucial in distributed systems to ensure that every node has an identical copy of the system's state.

### Contents

- **`system.go`**: Contains the main implementation of the distributed system, using the Paxos consensus algorithm.

## Features of the Distributed System Example

- **Node Representation**:
  - Each node in the system is represented as an independent instance, capable of proposing values and participating in the consensus process.
  - Nodes communicate with each other to agree on updates to the system state, ensuring that consistency is maintained.

- **Consensus Algorithms**:
  - The example uses Paxos and Raft to achieve distributed consensus, demonstrating how these algorithms help maintain consistency across nodes.
  - Each node can take on roles such as proposer, acceptor, and learner (in Paxos) or leader and follower (in Raft).

### Code Example

Below is an example of how to use the distributed system with Paxos as the consensus algorithm:

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

### How to Run the Distributed System Example

1. **Clone the Repository**:
   - Clone the repository and navigate to the `examples/distributed_system/` folder.

   ```bash
   git clone https://github.com/dkrizhanovskyi/consensus-algorithms-edu.git
   cd consensus-algorithms-edu/examples/distributed_system
   ```

2. **Build and Run the Code**:
   - Run the code using the Go compiler:

   ```bash
   go run system.go
   ```

3. **Output**:
   - The program will simulate consensus across multiple nodes, showing how each block is proposed, accepted, and committed to maintain consistency in the distributed system.

### Using Different Consensus Algorithms

This distributed system example can be adapted to use different consensus algorithms available in the `algorithms/` folder. Below is a description of how to modify the code to use Raft:

1. **Using Raft for Consensus**:
   - Import the `raft` package instead of `paxos`:

   ```go
   import "consensus-algorithms-edu/algorithms/raft"
   ```

   - Initialize the Raft network and use the leader to replicate data:

   ```go
   networkSize := 5
   blockchain := raft.NewRaftNetwork(networkSize)

   blockchain.Leader.Lead("First distributed system data")
   blockchain.Leader.Lead("Second distributed system data")
   blockchain.Leader.Lead("Third distributed system data")
   ```

### Key Concepts Demonstrated

- **Distributed Consensus in Action**: This example shows how distributed consensus ensures consistency across nodes, even in scenarios where some nodes may fail or the network experiences partitioning.
- **Role Differentiation**: In Paxos, nodes can be proposers, acceptors, or learners, while in Raft, nodes take on the roles of leader, follower, or candidate. These roles ensure that consensus is coordinated effectively in the distributed system.
- **Fault Tolerance**: The implementation demonstrates how distributed consensus algorithms tolerate failures, allowing the system to continue functioning correctly even if some nodes crash or behave erratically.

## Advantages of This Distributed System Example

- **Modular Design**: The example is designed to be modular, allowing different consensus algorithms to be easily swapped in to see how they affect the behavior of the distributed system.
- **Educational Value**: It serves as a great learning tool for understanding the principles of distributed consensus, fault tolerance, and how consensus is maintained among multiple nodes.
- **Real-World Application**: Concepts covered here are fundamental to real-world distributed systems like distributed databases, financial ledgers, and configuration management tools.

## Limitations

- **Simplified Environment**: This implementation simulates a distributed system on a single machine, abstracting away the complexities of network latency, partitioning, and message loss that occur in real distributed environments.
- **Single Threaded**: Each node is represented as a function in a single program, rather than as independent, concurrently running entities. In real distributed systems, each node runs independently on separate machines or processes.

## Practical Use Cases

- **Distributed Databases**: Systems like Google Spanner and etcd use consensus protocols like Paxos and Raft to maintain consistent state across geographically distributed nodes.
- **Cluster Management**: Tools like Consul and Kubernetes use Raft to manage cluster state and leader election, ensuring that updates are reliably propagated across the cluster.
- **Key-Value Stores**: Consensus algorithms like Paxos and Raft are also used to replicate key-value stores, ensuring that all replicas remain consistent in the face of node failures or network issues.

## Conclusion

This distributed system example demonstrates how consensus algorithms like Paxos and Raft are used to maintain a consistent state across nodes in a distributed environment. By experimenting with these consensus mechanisms, users can better understand how distributed systems ensure fault tolerance, reliability, and data consistency, even when nodes fail or network partitions occur.

### License

This documentation and the associated code are licensed under the MIT License.