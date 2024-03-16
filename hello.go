package main

import (
	"fmt"
	"strconv"
)

// Account struct represents a bank account
type Account struct {
	accountNumber int
	balance       float64
}

// CreateAccount creates a new account with the given account number and initial balance
func CreateAccount(accountNumber int, initialBalance float64) *Account {
	return &Account{accountNumber, initialBalance}
}

// Deposit adds the specified amount to the account balance
func (acc *Account) Deposit(amount float64) {
	acc.balance += amount
	fmt.Printf("Deposited %.2f. New balance: %.2f\n", amount, acc.balance)
}

// Withdraw subtracts the specified amount from the account balance
func (acc *Account) Withdraw(amount float64) {
	if amount > acc.balance {
		fmt.Println("Insufficient funds")
		return
	}
	acc.balance -= amount
	fmt.Printf("Withdrawn %.2f. New balance: %.2f\n", amount, acc.balance)
}

// Balance returns the current balance of the account
func (acc *Account) Balance() float64 {
	return acc.balance
}

func main() {
	// Create a new account with an initial balance of 0
	account := CreateAccount(123456789, 0.0)

	// Prompt the user to input transactions until they enter "done"
	for {
		var input string
		fmt.Print("Enter transaction amount (or 'done' to finish): ")
		fmt.Scanln(&input)

		if input == "done" {
			break
		}

		amount, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}

		var transactionType string
		fmt.Print("Enter transaction type (deposit/withdraw): ")
		fmt.Scanln(&transactionType)

		switch transactionType {
		case "deposit":
			account.Deposit(amount)
		case "withdraw":
			account.Withdraw(amount)
		default:
			fmt.Println("Invalid transaction type. Please enter 'deposit' or 'withdraw'.")
		}
	}

	// Check the final balance
	fmt.Printf("Final balance: %.2f\n", account.Balance())
}
