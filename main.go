
package main

import (
    "fmt"
    "net/http"
    "strconv"
    "text/template"
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

// Deposit adds the specified amount to the account balance and logs the transaction
func (acc *Account) Deposit(amount float64) {
    acc.balance += amount
    log := fmt.Sprintf("Deposited %.2f. New balance: %.2f", amount, acc.balance)
    addTransactionLog(log)
}

// Withdraw subtracts the specified amount from the account balance and logs the transaction
func (acc *Account) Withdraw(amount float64) {
    if amount > acc.balance {
        log := "Insufficient funds"
        addTransactionLog(log)
        return
    }
    acc.balance -= amount
    log := fmt.Sprintf("Withdrawn %.2f. New balance: %.2f", amount, acc.balance)
    addTransactionLog(log)
}

// Balance returns the current balance of the account
func (acc *Account) Balance() float64 {
    return acc.balance
}

var transactionLogs []string

// Function to add transaction logs
func addTransactionLog(log string) {
    transactionLogs = append(transactionLogs, log)
}

// Function to get transaction logs
func getTransactionLogs() []string {
    return transactionLogs
}

// ClearCacheHandler handles clearing the cache and resetting logs and balance
func ClearCacheHandler(w http.ResponseWriter, r *http.Request) {
    // Clear transaction logs
    transactionLogs = []string{}

    // Reset balance
    account.balance = 0.0

    // Write response
    fmt.Fprintf(w, "Cache cleared successfully. Logs and balance reset.")
}

var account *Account

func main() {
    // Create a new account with an initial balance of 0
    account = CreateAccount(123456789, 0.0)

    // Define the handler function for the home page
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        tmpl := template.Must(template.ParseFiles("index.html"))

        // Execute the template with the current balance and transaction logs
        data := struct {
            Balance float64
            Logs    []string
        }{
            Balance: account.Balance(),
            Logs:    getTransactionLogs(),
        }
        tmpl.Execute(w, data)
    })

    // Define the handler function for processing transactions
    http.HandleFunc("/transaction", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }

        // Parse form data
        if err := r.ParseForm(); err != nil {
            http.Error(w, "Failed to parse form data", http.StatusInternalServerError)
            return
        }

        amount, err := strconv.ParseFloat(r.Form.Get("amount"), 64)
        if err != nil {
            http.Error(w, "Invalid amount", http.StatusBadRequest)
            return
        }

        transactionType := r.Form.Get("transactionType")
        switch transactionType {
        case "deposit":
            account.Deposit(amount)
        case "withdraw":
            account.Withdraw(amount)
        default:
            http.Error(w, "Invalid transaction type", http.StatusBadRequest)
            return
        }

        // Redirect back to the home page after processing the transaction
        http.Redirect(w, r, "/", http.StatusSeeOther)
    })

    // Define the handler function for clearing the cache
    http.HandleFunc("/clearcache", ClearCacheHandler)

    // Start the server
    fmt.Println("Server running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}