package assignment01bca

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// Block represents a basic block in a blockchain.
type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	Hash         string
}

// Blockchain represents a chain of blocks.
type Blockchain struct {
	Blocks []*Block
}

// NewBlock creates a new block with the given transaction, nonce, and previous hash.
func NewBlock(transaction string, nonce int, previousHash string) *Block {
	block := &Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
	}
	block.Hash = CalculateHash(fmt.Sprintf("%s%d%s", block.Transaction, block.Nonce, block.PreviousHash))
	return block
}

// DisplayBlocks prints all the blocks in a nice format.
func (bc *Blockchain) DisplayBlocks() {
	for i, block := range bc.Blocks {
		fmt.Printf("Block %d:\n", i)
		fmt.Printf("  Transaction: %s\n", block.Transaction)
		fmt.Printf("  Nonce: %d\n", block.Nonce)
		fmt.Printf("  Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("  Block Hash: %s\n\n", block.Hash)
	}
}

// ChangeBlock updates the transaction of a given block reference.
func (bc *Blockchain) ChangeBlock(blockIndex int, newTransaction string) {
	if blockIndex >= 0 && blockIndex < len(bc.Blocks) {
		bc.Blocks[blockIndex].Transaction = newTransaction
		bc.Blocks[blockIndex].Hash = CalculateHash(fmt.Sprintf("%s%d%s", newTransaction, bc.Blocks[blockIndex].Nonce, bc.Blocks[blockIndex].PreviousHash))
	}
}

// VerifyChain verifies the integrity of the blockchain.
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

// func assignment01bca() {
// 	blockchain := &Blockchain{Blocks: []*Block{}}

// 	fmt.Println("Simple Blockchain Menu")
// 	fmt.Println("----------------------")

// 	for {
// 		fmt.Println("1. New Block")
// 		fmt.Println("2. Display Blocks")
// 		fmt.Println("3. Change Block Transaction")
// 		fmt.Println("4. Verify Blockchain")
// 		fmt.Println("5. Exit")

// 		var choice int
// 		fmt.Print("Enter your choice: ")
// 		_, err := fmt.Scan(&choice)
// 		if err != nil {
// 			fmt.Println("Invalid input. Please enter a valid choice.")
// 			continue
// 		}

// 		switch choice {
// 		case 1:
// 			fmt.Print("Enter transaction: ")
// 			var transaction string
// 			_, _ = fmt.Scan(&transaction)

// 			fmt.Print("Enter nonce (an integer): ")
// 			var nonce int
// 			_, err := fmt.Scan(&nonce)
// 			if err != nil {
// 				fmt.Println("Invalid nonce. Please enter a valid integer.")
// 				continue
// 			}

// 			previousHash := ""
// 			if len(blockchain.Blocks) > 0 {
// 				previousHash = blockchain.Blocks[len(blockchain.Blocks)-1].Hash
// 			}

// 			blockchain.Blocks = append(blockchain.Blocks, NewBlock(transaction, nonce, previousHash))
// 			fmt.Println("Block added successfully!")

// 		case 2:
// 			fmt.Println("Blockchain:")
// 			blockchain.DisplayBlocks()

// 		case 3:
// 			fmt.Print("Enter the block index to change transaction: ")
// 			var indexStr string
// 			_, _ = fmt.Scan(&indexStr)
// 			index, err := strconv.Atoi(indexStr)
// 			if err != nil || index < 0 || index >= len(blockchain.Blocks) {
// 				fmt.Println("Invalid block index. Please enter a valid integer.")
// 				continue
// 			}
// 			fmt.Print("Enter new transaction: ")
// 			var newTransaction string
// 			_, _ = fmt.Scan(&newTransaction)
// 			blockchain.ChangeBlock(index, newTransaction)
// 			fmt.Println("Transaction changed successfully!")

// 		case 4:
// 			if blockchain.VerifyChain() {
// 				fmt.Println("Blockchain is valid.")
// 			} else {
// 				fmt.Println("Blockchain is invalid.")
// 			}

// 		case 5:
// 			fmt.Println("Exiting...")
// 			return

// 		default:
// 			fmt.Println("Invalid choice. Please select a valid option.")
// 		}
// 	}
// }
