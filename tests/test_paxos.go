package tests

import (
    "testing"
    "consensus-algorithms-edu/algorithms/paxos"
)

func TestPaxos(t *testing.T) {
    blockchain := paxos.NewPaxosNetwork(5)

    blockchain.RunPaxos("Test block 1", 1)
    blockchain.RunPaxos("Test block 2", 2)

    if len(blockchain.Blocks) != 3 {
        t.Errorf("Expected 3 blocks, got %d", len(blockchain.Blocks))
    }

    lastBlock := blockchain.Blocks[len(blockchain.Blocks)-1]
    if lastBlock.Data != "Test block 2" {
        t.Errorf("Expected 'Test block 2', got '%s'", lastBlock.Data)
    }
}
