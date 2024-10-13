# Consensus Algorithms Educational Repository

Author: **D. Krizhanovskyi**  
**Applied Cryptographer and Blockchain Architect**

## About Me

I am an applied cryptographer and blockchain architect with extensive experience in the blockchain space. My journey includes significant contributions to the **Solana ecosystem** and a deep understanding of **peer-to-peer network systems** and **distributed system architecture**. My expertise is focused on designing secure, scalable blockchain protocols that address real-world challenges in the areas of decentralization, security, and performance optimization.

### Academic Background

- **Master of Science in Applied Mathematics**  
  *Lviv Polytechnic National University (2020-2022)*  
  *Thesis*: "Mathematical Modeling and Computational Techniques in Modern Cryptography."

- **Master of Science in Information Systems and Technologies**  
  *Odessa National Polytechnic University (2019-2021)*  
  *Thesis*: "Development and Implementation of Secure Information Systems in Modern Cryptography."

## Overview of This Repository

This repository is an educational resource focused on **consensus algorithms** for blockchain and distributed systems. It is intended for students, developers, and blockchain enthusiasts who want to understand the principles behind consensus and the technical implementations that keep decentralized systems secure and consistent.

The content covers a variety of **consensus mechanisms**, each implemented in the **Go programming language** for its simplicity, efficiency, and suitability for building scalable, high-performance blockchain applications. The repository includes the following key consensus mechanisms:

### Included Consensus Algorithms

1. **Proof of Work (PoW)**:
   - A battle-tested consensus mechanism used by **Bitcoin**. It secures the blockchain through computational puzzles that miners solve to create blocks.
2. **Proof of Stake (PoS)**:
   - Used in networks like **Ethereum 2.0**. It leverages validators who stake their tokens to validate transactions, making it more energy-efficient compared to PoW.
3. **Delegated Proof of Stake (DPoS)**:
   - An improved version of PoS that relies on a voting mechanism to elect delegates, providing fast transaction processing while maintaining decentralization. **EOS** and **TRON** are examples of networks using DPoS.
4. **Practical Byzantine Fault Tolerance (PBFT)**:
   - A consensus mechanism used in **permissioned blockchains** like **Hyperledger Fabric**. It can tolerate Byzantine faults, allowing a subset of nodes to act arbitrarily or even maliciously.
5. **Raft**:
   - A leader-based consensus algorithm often used for **log replication** and **distributed database consistency**, with a strong focus on simplicity. Tools like **etcd** and **Consul** use Raft.
6. **Paxos**:
   - A consensus algorithm used in **distributed systems** for replicated logs and state machines. It is well-known for its mathematical robustness and its use in systems like **Google Spanner**.

### Structure of This Repository

The repository is organized as follows:

- **algorithms/**: Contains the implementation of each consensus algorithm.
  - **pow/**: Implementation of Proof of Work.
  - **pos/**: Implementation of Proof of Stake.
  - **dpos/**: Implementation of Delegated Proof of Stake.
  - **pbft/**: Implementation of Practical Byzantine Fault Tolerance.
  - **raft/**: Implementation of Raft consensus.
  - **paxos/**: Implementation of Paxos.
  
- **examples/**: Practical examples showcasing different consensus algorithms in action.
  - **blockchain_example/**: A basic blockchain implementation demonstrating consensus.
  - **distributed_system/**: Demonstrates how consensus mechanisms maintain consistency in distributed environments.
  - **voting_example/**: A voting system example using the DPoS consensus mechanism.
  
- **docs/**: Contains detailed documentation about each consensus algorithm.
  - **PoW.md**: Overview of Proof of Work.
  - **PoS.md**: Overview of Proof of Stake.
  - **DPoS.md**: Overview of Delegated Proof of Stake.
  - **PBFT.md**: Overview of Practical Byzantine Fault Tolerance.
  - **Raft.md**: Overview of Raft.
  - **Paxos.md**: Overview of Paxos.
  - **shared_scenarios.md**: Shared scenarios and use cases comparing different consensus algorithms.

### Key Features of the Repository

- **Hands-On Code**: Each consensus mechanism is implemented in Go, providing easy-to-follow code examples.
- **Documentation**: Each algorithm is paired with detailed documentation that covers its purpose, how it works, and example use cases.
- **Practical Scenarios**: Shared scenarios and examples are included to highlight real-world use cases for each consensus mechanism, making it easier to understand when and why to use each type of consensus.

## Intended Audience

- **Beginner Developers**: Who are learning about blockchain technology and want to get hands-on with consensus algorithms.
- **Experienced Engineers**: Who are interested in exploring different consensus mechanisms and comparing their effectiveness.
- **Students**: Who are studying distributed systems and want to understand the theory and implementation of consensus algorithms.

## Getting Started

To start exploring the consensus algorithms:

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/dkrizhanovskyi/consensus-algorithms-edu.git
   cd consensus-algorithms-edu
   ```

2. **Run Examples**:

   Navigate to the `examples/` folder and run one of the provided examples to see the consensus algorithm in action:

   ```bash
   cd examples/blockchain_example
   go run blockchain.go
   ```

3. **Explore Algorithms**:

   Check out the `algorithms/` directory to explore the source code for each consensus mechanism and understand their individual characteristics.

## Why Use This Repository?

- **Modular and Extendable**: Each consensus algorithm is implemented in a modular way, allowing you to swap out different consensus mechanisms easily.
- **Educational Focus**: The repository is designed to help people learn, featuring clear examples, detailed comments, and structured documentation.
- **Real-World Insight**: Gain an understanding of how blockchain systems work in real-world applications by exploring various consensus protocols used in modern decentralized networks.

## Future Enhancements

- **Add More Consensus Mechanisms**: Expand the repository to include newer consensus mechanisms like **Tendermint** and **HoneyBadger BFT**.
- **More Real-World Examples**: Add examples showing the integration of these consensus algorithms into larger distributed systems.
- **Interactive Tutorials**: Provide interactive examples using Docker to simulate a distributed environment for testing and learning purposes.

## License

This repository is licensed under the **MIT License**, making it free and open to use and extend for educational purposes.

---

If you have any questions or would like to contribute, feel free to reach out or submit a pull request. Happy coding!

