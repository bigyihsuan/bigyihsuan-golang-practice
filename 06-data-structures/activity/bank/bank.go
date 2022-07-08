package main

import "fmt"

type EntityType string
type AccountType string

const (
	individual EntityType = "individual"
	business              = "business"
)

const (
	checking   AccountType = "checking"
	savings                = "savings"
	investment             = "investment"
)

type Entity struct {
	Id      int
	Address string
	Type    EntityType // individual or business
}

type Account struct {
	Number       string
	Owner        Entity
	Balance      float64
	InterestRate float64
	Type         AccountType // checking, savings, investment
}

type Wallet struct {
	Id       string
	Accounts []Account
	Owner    Entity
}

func (a *Account) Withdraw(amount float64) {
	if a.CanWithdraw(amount) {
		a.Balance -= amount
	}
}

func (a *Account) CanWithdraw(amount float64) bool {
	return a.Balance > amount && a.Balance >= 0
}

func (a *Account) Deposit(amount float64) {
	a.Balance += amount
}

func (a *Account) ApplyInterest() {
	var apr float64
	switch a.Owner.Type {
	case individual:
		switch a.Type {
		case checking:
			apr = 0.01
		case savings:
			apr = 0.02
		case investment:
			apr = 0.05
		}
	case business:
		switch a.Type {
		case checking:
			apr = 0.005
		case savings:
			apr = 0.01
		case investment:
			apr = 0.02
		}
	}
	a.Balance += a.Balance * apr
}

func (source *Account) wireTo(dest *Account, amount float64) {
	if source.CanWithdraw(amount) {
		source.Withdraw(amount)
		dest.Deposit(amount)
	}
}

func (e *Entity) ChangeAddress(newAddress string) {
	e.Address = newAddress
}

func (w Wallet) GetAccountByType(t AccountType) []Account {
	var out []Account
	for _, account := range w.Accounts {
		if account.Type == t {
			out = append(out, account)
		}
	}
	return out
}

func (w Wallet) DisplayAccounts() {
	check := w.GetAccountByType(checking)
	invest := w.GetAccountByType(investment)
	save := w.GetAccountByType(savings)
	fmt.Printf("%v's accounts: ", w.Owner.Id)
	fmt.Printf("checking accounts: %v\n", check)
	fmt.Printf("investment accounts: %v\n", invest)
	fmt.Printf("savings accounts: %v\n", save)
}

func (w Wallet) Balance() float64 {
	var total float64
	for _, account := range w.Accounts {
		total += account.Balance
	}
	return total
}

func (w *Wallet) Wire(source Account, destination Account, amount float64) {
	if !source.CanWithdraw(amount) {
		fmt.Printf("ERROR: not enough balance in account %v to withdraw $%v\n", source.Number, amount)
	}
	source.wireTo(&destination, amount)
}

func main() {
	bob := Entity{
		123, "123 main", individual,
	}
	pat := Entity{
		567, "broadway", business,
	}

	a1 := Account{
		"1", bob, 555.55, 0.1, checking,
	}
	a2 := Account{
		"2", bob, 44444.44, 0.2, investment,
	}
	a3 := Account{
		"3", pat, 123.12, 0.1, savings,
	}
	a4 := Account{
		"4", pat, 456.45, 0.1, checking,
	}

	bobWallet := Wallet{
		"bob's wallet", []Account{a1, a2}, bob,
	}

	patWallet := Wallet{
		"pat's wallet", []Account{a3, a4}, pat,
	}

	bobWallet.DisplayAccounts()
	patWallet.DisplayAccounts()
	
	bobWallet.Wire(a1, a2, 10)
	bobWallet.Wire(a1, a2, -10)
	bobWallet.Wire(a1, a2, a1.Balance*10)
	bobWallet.Wire(a1, a4, 300)
	a4.ApplyInterest()
	a4.ApplyInterest()
	a4.ApplyInterest()
	a4.ApplyInterest()
	bobWallet.DisplayAccounts()
	patWallet.DisplayAccounts()
}
