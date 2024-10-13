package tests

import (
    "testing"
    "consensus-algorithms-edu/algorithms/raft"
)

func TestRaft(t *testing.T) {
    blockchain := raft.NewRaftNetwork(5)

    blockchain.Leader.Lead("Test block 1")
    blockchain.Leader.Lead("Test block 2")

    if len(blockchain.Blocks) != 3 {
        t.Errorf("Expected 3 blocks, got %d", len(blockchain.Blocks))
    }

    lastBlock := blockchain.Blocks[len(blockchain.Blocks)-1]
    if lastBlock.Data != "Test block 2" {
        t.Errorf("Expected 'Test block 2', got '%s'", lastBlock.Data)
    }
}
