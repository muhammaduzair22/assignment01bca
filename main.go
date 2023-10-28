// Muhammad Uzair 20i-2341
package main

import (
	// "20i-2341_A/assignment01bca"
	"fmt"
	"strconv"

	"github.com/muhammaduzair22/assignment01bca"
)

// github usernamer: muhammaduzair22
// github.com/muhammaduzair22/assignment01bca

func main() {
	assignment01bca := &assignment01bca.Blockchain{Blocks: []*assignment01bca.Block{}}

	fmt.Println("Blockchain Menu")
	fmt.Println("----------------------")

	for {
		fmt.Println("1. New Block")
		fmt.Println("2. Display Blocks")
		fmt.Println("3. Change Block")
		fmt.Println("4. Verify Chain")
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
			fmt.Println("----------------------")

		case 2:
			fmt.Println("Blockchain:")
			assignment01bca.DisplayBlocks()
			fmt.Println("----------------------")

		case 3:
			fmt.Print("Enter the block index to change transaction: ")
			var indexStr string
			_, _ = fmt.Scan(&indexStr)
			index, err := strconv.Atoi(indexStr)
			if err != nil || index < 0 || index >= len(assignment01bca.Blocks) {
				fmt.Println("----------------------")
				fmt.Println("Invalid block index. Please enter a valid integer.")
				fmt.Println("----------------------")
				continue
			}
			fmt.Println("----------------------")
			fmt.Print("Enter new transaction: ")
			var newTransaction string
			_, _ = fmt.Scan(&newTransaction)
			assignment01bca.ChangeBlock(index, newTransaction)
			fmt.Println("Transaction changed successfully!")
			fmt.Println("----------------------")

		case 4:
			if assignment01bca.VerifyChain() {
				fmt.Println("Blockchain is valid.")
				fmt.Println("----------------------")

			} else {
				fmt.Println("Blockchain is invalid.")
				fmt.Println("----------------------")

			}

		case 5:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice. Please select a valid option.")
			fmt.Println("----------------------")

		}
	}
}
