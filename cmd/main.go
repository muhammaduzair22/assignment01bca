package main

import (
	"fmt"
	"strconv"

	"github.com/muhammaduzair22/assignment01bca" // Import the blockchain package
)

func main() {
	assignment01bca := &assignment01bca.assignment01bca{Blocks: []*assignment01bca.Block{}}

	fmt.Println("Simple Blockchain Menu")
	fmt.Println("----------------------")

	for {
		fmt.Println("1. Add Block")
		fmt.Println("2. Display Blocks")
		fmt.Println("3. Change Block Transaction")
		fmt.Println("4. Verify Blockchain")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid choice.")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Enter transaction: ")
			var transaction string
			_, _ = fmt.Scan(&transaction)

			fmt.Print("Enter nonce (an integer): ")
			var nonce int
			_, err := fmt.Scan(&nonce)
			if err != nil {
				fmt.Println("Invalid nonce. Please enter a valid integer.")
				continue
			}

			previousHash := ""
			if len(assignment01bca.Blocks) > 0 {
				previousHash = assignment01bca.Blocks[len(assignment01bca.Blocks)-1].Hash
			}

			assignment01bca.Blocks = append(assignment01bca.Blocks, assignment01bca.NewBlock(transaction, nonce, previousHash))
			fmt.Println("Block added successfully!")

		case 2:
			fmt.Println("Blockchain:")
			assignment01bca.DisplayBlocks()

		case 3:
			fmt.Print("Enter the block index to change transaction: ")
			var indexStr string
			_, _ = fmt.Scan(&indexStr)
			index, err := strconv.Atoi(indexStr)
			if err != nil || index < 0 || index >= len(assignment01bca.Blocks) {
				fmt.Println("Invalid block index. Please enter a valid integer.")
				continue
			}
			fmt.Print("Enter new transaction: ")
			var newTransaction string
			_, _ = fmt.Scan(&newTransaction)
			assignment01bca.ChangeBlock(index, newTransaction)
			fmt.Println("Transaction changed successfully!")

		case 4:
			if assignment01bca.VerifyChain() {
				fmt.Println("Blockchain is valid.")
			} else {
				fmt.Println("Blockchain is invalid.")
			}

		case 5:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}
