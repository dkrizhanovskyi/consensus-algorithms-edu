package tests

import (
    "testing"
    "consensus-algorithms-edu/algorithms/dpos"
)

func TestDPoS(t *testing.T) {
    delegates := []string{"Alice", "Bob", "Charlie"}
    voters := map[string]string{}

    blockchain := dpos.NewBlockchain(delegates, voters)

    blockchain.Vote("Voter1", "Alice")
    blockchain.Vote("Voter2", "Bob")
    blockchain.CountVotes()

    blockchain.AddBlock("Test block 1")
    blockchain.AddBlock("Test block 2")

    if len(blockchain.Blocks) != 3 {
        t.Errorf("Expected 3 blocks, got %d", len(blockchain.Blocks))
    }

    lastBlock := blockchain.Blocks[len(blockchain.Blocks)-1]
    if lastBlock.Data != "Test block 2" {
        t.Errorf("Expected 'Test block 2', got '%s'", lastBlock.Data)
    }
}
