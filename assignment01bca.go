package assignment01bca

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	Hash         string
}

type Blockchain struct {
	Blocks []*Block
}

func NewBlock(transaction string, nonce int, previousHash string) *Block {
	block := &Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
	}
	block.Hash = CalculateHash(fmt.Sprintf("%s%d%s", block.Transaction, block.Nonce, block.PreviousHash))
	return block
}

func (bc *Blockchain) DisplayBlocks() {
	for i, block := range bc.Blocks {
		fmt.Printf("Block %d:\n", i)
		fmt.Printf("  Transaction: %s\n", block.Transaction)
		fmt.Printf("  Nonce: %d\n", block.Nonce)
		fmt.Printf("  Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("  Block Hash: %s\n\n", block.Hash)
	}
}

func (bc *Blockchain) ChangeBlock(blockIndex int, newTransaction string) {
	if blockIndex >= 0 && blockIndex < len(bc.Blocks) {
		bc.Blocks[blockIndex].Transaction = newTransaction
		bc.Blocks[blockIndex].Hash = CalculateHash(fmt.Sprintf("%s%d%s", newTransaction, bc.Blocks[blockIndex].Nonce, bc.Blocks[blockIndex].PreviousHash))
	}
}

func (bc *Blockchain) VerifyChain() bool {
	for i := 1; i < len(bc.Blocks); i++ {
		currentBlock := bc.Blocks[i]
		previousBlock := bc.Blocks[i-1]

		if currentBlock.Hash != CalculateHash(fmt.Sprintf("%s%d%s", currentBlock.Transaction, currentBlock.Nonce, currentBlock.PreviousHash)) {
			return false
		}

		if currentBlock.PreviousHash != previousBlock.Hash {
			return false
		}
	}

	return true
}

// CalculateHash calculates the hash of a given string.
func CalculateHash(stringToHash string) string {
	hashBytes := sha256.Sum256([]byte(stringToHash))
	return hex.EncodeToString(hashBytes[:])
}
