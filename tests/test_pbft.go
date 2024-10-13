package tests

import (
    "testing"
    "consensus-algorithms-edu/algorithms/pbft"
)

func TestPBFT(t *testing.T) {
    blockchain := pbft.NewPBFTNetwork(5)

    blockchain.RunPBFT("Test block 1")
    blockchain.RunPBFT("Test block 2")

    if len(blockchain.Blocks) != 3 {
        t.Errorf("Expected 3 blocks, got %d", len(blockchain.Blocks))
    }

    lastBlock := blockchain.Blocks[len(blockchain.Blocks)-1]
    if lastBlock.Data != "Test block 2" {
        t.Errorf("Expected 'Test block 2', got '%s'", lastBlock.Data)
    }
}
