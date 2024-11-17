# Blockchain Based Voting Network

Welcome to the **Blockchain Voting Network** project!

Blockchain-Based Voting System (BBVS) leverages the unique properties of blockchain technology. Leveraging the Hyperledger Fabric framework and smart contracts, this system replaces a central voting database with a private, permissioned blockchain that records vote on an immutable ledger, ensuring tamper-proof data integrity. Blockchain’s core characteristics - transparency, immutability, and accountability - underscore its potential to secure elections effectively.

## Network Architecture

The BBVS prototype was developed with three operational validators (referred to as "Organizations") configured as a unified organization for Proof of Concept (PoC) purposes. Among the available consensus algorithms in Hyperledger Fabric - RAFT, Kafka, Solo, and PBFT. RAFT were chosen for this implementation. RAFT, a crash-fault-tolerant (CFT) ordering service based on a leader-follower model, offers a simplified setup and administration as well as a more decentralized approach to consensus management compared to Kafka, Solo, and PBFT. The network is structured with three organizations, designated Org1, Org2, and Org3, each containing two peers and an ordering node. Each organization is managed by a single administrator by default. All organizations are authorized to create channels through the ordering service and are part of the consortium. The network is designed with two channels: one channel logs the transactions for voters’ ballots, and the other channel stores mapping data between the voting token.

Hyperledger Fabric leverages cryptographic materials and Docker to create secure containerized services, supporting a modular, scalable, and permissioned blockchain network. This framework enables identity management, access control, and configurable consensus, making it ideal for applications that require privacy and data integrity.

![image](https://github.com/user-attachments/assets/0efbc52b-18db-41a4-b480-eb86f78fe7b2)
<p align="center">
    BlockChain Network Arcitecture
</p>

<!--
## To Get Started

1. **Clone the Repository**  
   To start, clone this repository to your local machine using the following command:  
   ```bash
   git clone https://github.com/satyam-thakur/Blockchain_Voting_Network.git

2. Run the composite bash Script ./deployblockchain.sh
   ```bash
   cd Blockchain_Voting_Network
   ./deployblockchain.sh



