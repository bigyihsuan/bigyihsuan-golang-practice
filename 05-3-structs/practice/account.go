package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Transaction struct {
	account_id int
	time       int64
	status     bool // false if failed
	amount     float64
}
type BankTransactions []Transaction

func newTransaction(account_id int, status bool, amount float64) Transaction {
	return Transaction{
		account_id: account_id,
		time:       time.Now().Unix(),
		status:     status,
		amount:     amount,
	}
}

type Account struct {
	name    string
	id      int
	address string
	balance float64
}

func (a *Account) deposit(amount float64, transactionLog BankTransactions) BankTransactions {
	if amount < 0 {
		fmt.Printf("ERR: deposit amount (%v) is negative, aborted\n", amount)
		transactionLog = append(transactionLog, newTransaction(a.id, false, amount))
	} else {
		transactionLog = append(transactionLog, newTransaction(a.id, true, amount))
		a.balance += amount
	}
	return transactionLog
}

func (a *Account) withdraw(amount float64, transactionLog BankTransactions) BankTransactions {
	if amount < 0 {
		fmt.Printf("ERR: withdraw amount (%v) is negative, aborted\n", amount)
		transactionLog = append(transactionLog, newTransaction(a.id, false, -amount))
	} else if a.balance < amount {
		fmt.Printf("ERR: withdraw amount (%v) is more than account balance (%v), aborted\n", amount, a.balance)
		transactionLog = append(transactionLog, newTransaction(a.id, false, -amount))
	} else {
		transactionLog = append(transactionLog, newTransaction(a.id, true, -amount))
		a.balance -= amount
	}
	return transactionLog
}

type Bank []*Account

func (b Bank) findAccountById(id int) *Account {
	if id >= len(b) {
		return nil
	}
	return b[id]
}

func (b Bank) sortByBalance() Bank {
	sort.Slice(b, func(i, j int) bool {
		return b[i].balance < b[j].balance
	})
	return b
}

func (b Bank) sortByName() Bank {
	sort.Slice(b, func(i, j int) bool {
		firstName1 := strings.Split(b[i].name, " ")[0]
		firstName2 := strings.Split(b[j].name, " ")[0]
		return firstName1 < firstName2
	})
	return b
}

func main() {
	var bank Bank
	var bankLog BankTransactions

	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < 3; i++ {
		account := new(Account)
		fmt.Print("Enter name: ")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)
		fmt.Print("Enter balance: ")
		balance, _ := reader.ReadString('\n')
		balance = strings.TrimSpace(balance)
		// fmt.Print("Enter address: ")
		// address, _ := reader.ReadString('\n')
		// address = strings.TrimSpace(address)

		account.id = i
		account.name = name
		account.balance, _ = strconv.ParseFloat(balance, 64)
		// account.address = address
		bank = append(bank, account)
		fmt.Println()
	}
	for i, e := range bank {
		fmt.Printf("before: %v", *e)
		bankLog = e.deposit(100+float64(i)*0.01, bankLog)
		bankLog = e.withdraw(float64(i)*25, bankLog)
		fmt.Printf("    after: %v", *e)
		fmt.Println()
	}

	for _, e := range bank {
		bankLog = e.deposit(-e.balance, bankLog)
		fmt.Println(bankLog)
	}

	// fmt.Println("\nsorted by balance:")
	// for _, e := range bank.sortByBalance() {
	// 	fmt.Println(e)
	// }

	// fmt.Println("\nsorted by first name:")
	// for _, e := range bank.sortByName() {
	// 	fmt.Println(e)
	// }
}
