# Blockchain Example Using Consensus Algorithms

This folder provides an example of a simple blockchain implementation that demonstrates how different consensus algorithms—such as Proof of Work (PoW), Proof of Stake (PoS), Delegated Proof of Stake (DPoS), Practical Byzantine Fault Tolerance (PBFT), Raft, and Paxos—can be used to maintain a decentralized ledger in a distributed network. The goal is to give an educational and practical demonstration of how blockchain works in combination with these consensus mechanisms.

## Overview

The blockchain implementation provided here is a simple ledger where blocks are appended to a chain after achieving consensus. Depending on the consensus algorithm used, the process of selecting which node adds the next block varies, providing an excellent way to compare the different approaches.

### Contents

- **`blockchain.go`**: Contains the main implementation of the blockchain example, utilizing a specific consensus algorithm (initially configured for Proof of Work).

## Features of the Blockchain Example

- **Block Structure**:
  - Each block contains an **index**, a **timestamp**, a **data** field, the **hash** of the previous block, and a **nonce** or **validator** depending on the consensus algorithm.
  - The blockchain is initialized with a **genesis block**, and additional blocks are added by the consensus algorithm in use.

- **Consensus Algorithms**:
  - This example can be adapted to use different consensus algorithms to add blocks to the blockchain. Each algorithm introduces different methods for deciding how blocks are added, ensuring network security and consistency.

### Code Example

Below is an example of how to use the blockchain with Proof of Work as the consensus algorithm:

```go
package main

import (
    "fmt"
    "consensus-algorithms-edu/algorithms/pow"
)

func main() {
    blockchain := pow.NewBlockchain()

    blockchain.AddBlock("First block data")
    blockchain.AddBlock("Second block data")
    blockchain.AddBlock("Third block data")

    for _, block := range blockchain.Blocks {
        fmt.Printf("Index: %d\nTimestamp: %s\nData: %s\nPrevious Hash: %s\nHash: %s\nNonce: %d\n\n", 
            block.Index, block.Timestamp, block.Data, block.PrevHash, block.Hash, block.Nonce)
    }
}
```

### How to Run the Blockchain Example

1. **Clone the Repository**:
   - Clone the repository and navigate to the `examples/blockchain_example/` folder.

   ```bash
   git clone https://github.com/dkrizhanovsyi/consensus-algorithms-edu.git
   cd consensus-algorithms-edu/examples/blockchain_example
   ```

2. **Build and Run the Code**:
   - Run the code using the Go compiler:

   ```bash
   go run blockchain.go
   ```

3. **Output**:
   - The program will print out the details of each block added to the blockchain, showing the **index**, **timestamp**, **data**, **previous hash**, **current hash**, and **nonce** for each block.

### Using Different Consensus Algorithms

This blockchain example can be adapted to use other consensus algorithms available in the `algorithms/` folder. Below is a description of how to modify the code to use a different consensus mechanism:

1. **Proof of Stake (PoS)**:
   - Import the `pos` package instead of `pow`:

   ```go
   import "consensus-algorithms-edu/algorithms/pos"
   ```

   - Initialize the blockchain with validators and stakes:

   ```go
   validators := []string{"Alice", "Bob"}
   stakes := map[string]int{
       "Alice": 60,
       "Bob":   40,
   }

   blockchain := pos.NewBlockchain(validators, stakes)
   ```

2. **Delegated Proof of Stake (DPoS)**:
   - Import the `dpos` package and initialize the blockchain with a list of delegates:

   ```go
   import "consensus-algorithms-edu/algorithms/dpos"

   delegates := []string{"Alice", "Bob", "Charlie"}
   voters := map[string]string{}

   blockchain := dpos.NewBlockchain(delegates, voters)
   blockchain.Vote("Voter1", "Alice")
   blockchain.CountVotes()
   ```

3. **PBFT, Raft, and Paxos**:
   - Similarly, you can use the `pbft`, `raft`, or `paxos` packages to initialize a blockchain with multiple nodes participating in consensus:

   ```go
   import "consensus-algorithms-edu/algorithms/pbft"

   blockchain := pbft.NewPBFTNetwork(5)
   blockchain.RunPBFT("Transaction data")
   ```

### Key Concepts Demonstrated

- **Blockchain and Consensus Integration**: This example illustrates how different consensus algorithms can be integrated with a blockchain to achieve distributed consensus in different ways.
- **Security Through Consensus**: Each algorithm provides unique mechanisms to ensure that only valid transactions are added to the blockchain, providing fault tolerance and resilience against attacks.
- **Practical Comparison**: By running the same blockchain with different consensus algorithms, users can practically understand the differences in efficiency, security, and network overhead between each consensus method.

## Advantages of This Blockchain Example

- **Modularity**: The blockchain implementation is modular, allowing different consensus algorithms to be easily swapped in and out. This highlights the key differences in how each algorithm affects block creation.
- **Educational Value**: By experimenting with different consensus mechanisms, developers and students can gain a deeper understanding of the advantages and limitations of each approach in a practical setting.
- **Simplified Demonstration**: The example abstracts away the complexities of a real-world blockchain, allowing users to focus on understanding consensus mechanisms without being overwhelmed by the intricacies of full-scale blockchain development.

## Limitations

- **Simplified Environment**: This implementation is simplified and does not include many features present in production blockchains, such as advanced security measures, network communication, or scalability features.
- **Single Machine Execution**: All nodes and consensus mechanisms are simulated on a single machine. Real-world consensus occurs in a distributed network with independent nodes that communicate over a network.

## Practical Use Cases

- **Educational Tools**: This blockchain example can be used in classroom settings to teach students about how consensus algorithms work in distributed ledgers.
- **Prototyping**: Developers can use this example to prototype different consensus models before implementing them in more complex systems.
- **Comparative Analysis**: This modular implementation allows easy comparison between the performance and security of different consensus algorithms, which is useful for researchers and developers.

## Conclusion

This blockchain example provides a hands-on way to understand how consensus algorithms are integrated with blockchain technology. By exploring different consensus mechanisms—such as PoW, PoS, DPoS, PBFT, Raft, and Paxos—users can gain practical insights into the strengths, trade-offs, and applications of each approach. Whether used for education, research, or prototyping, this example serves as a versatile tool for anyone interested in consensus algorithms and blockchain technology.

### License

This documentation and the associated code are licensed under the MIT License.